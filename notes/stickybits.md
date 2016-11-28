SUID, GUID, Sticky Bits
===========

# SUID/GUID Quirks

This will cover some of the quirks of linux file permissions that I often forget.

First note that if you are trying to use SUID or GUID on a filesystem you need to check if nosuid is set. You can easily test
for this by running `mount | grep /filesystem` which will show all the mount params for the given filesystem. If you see 
nosuid in this list using the suid and guid bits will be ignored. Here is an excerpt from the mount man page about the mount option.

```
nosuid

Do not allow set-user-identifier or set-group-identifier bits to take effect. 
(This seems safe, but is in fact rather unsafe if you have suidperl(1) installed.)
```

Another thing to note is that only binary files properly use SUID/GUID bits for security reasons. More on this
can be seen [here](http://unix.stackexchange.com/questions/130906/why-does-setuid-not-work).

# SUID/GUID Example

SUID/GUID bits allow you to set the user or group that the program will be run as regardless of the user who executes the program.
More history on it can be found [here](http://www.thegeekstuff.com/2013/02/sticky-bit/). 

The following will create a file that has the setUID bit set however since this only sets the setUID bit and
does not set the executable bit as well this does nothing and other users are not able to run this file as the
user who created the file [Reference](http://unix.stackexchange.com/questions/158448/what-is-the-purpose-of-setuid-enabled-with-no-executable-bit).

This was the example c program I used for seeing if SUID was being set how I expected.

```
// binary.c
#include <stdio.h>
#include <stdlib.h>
#include <sys/types.h>
#include <unistd.h>

int main()
{
  printf("uid: %d \n", geteuid());
  system("id -u");
  return 0;
}
```

Then I ran the following to set the SUID bit.

```
$ gcc binary.c -o binary
$ chmod u+s binary
$ ls -la
total 0
-rwSrw-r--  1 michael michael    0 nov 26 20:39 binary
```

The following will make the file executable and then you can test running the file as another user.

```
$ chmod +x file.sh
$ ls -la
$ ls -la
total 0
-rwsrw-r--  1 michael michael    0 nov 26 20:39 binary
```

Note the output when running the program as another user will show that the c call `geteuid` shows you the
expected uid of the user who created the file where `id -u` shows you your user id still. To clear this up
you will want to look at `man geteuid` and note the difference between uid and effective uid. This gist of 
it is that the effective uid decides what permissions you have and the uid is the owner of the process.
More can be read about this [here](http://stackoverflow.com/questions/32455684/difference-between-real-user-id-effective-user-id-and-saved-user-id). 

To enable you system calls to access the same files you will need to use the setreuid call before calling system.
Here is an example of using setreuid which would then allow all future system calls to access files on the system
in the same way as if you were the user who owned the binary.

```
if (setreuid(1000, 1000)) perror("setuid reported failure");
```

I won't go into GUID because it works exactly the same as the above except with group id instead of user.

# Stick Bit

The sticky bit is a permission bit that protects the files within a directory. If the directory has the 
sticky bit set, a file can be deleted only by the owner of the file, the owner of the directory, or by root.
This special permission prevents a user from deleting other users' files from public directories such as /tmp.

## Extra

1. http://www.cis.syr.edu/~wedu/Teaching/IntrCompSec/LectureNotes_New/Set_UID.pdf
2. http://stackoverflow.com/questions/24590735/setuid-does-not-work-for-non-root
