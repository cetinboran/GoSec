# GoSec
+ This project encompasses two of my own creations, namely, goarg and gojson.
    + [**goarg**](https://github.com/cetinboran/goarg): It effortlessly parses command line arguments, providing you with the essential input you need.
    + [**gojson**](https://github.com/cetinboran/gojson): With just a few lines of code, you can establish your own miniature database using JSON.
+ You can also contribute to these projects and show your support.

## What is This?
+ This project is a locally running password management application. It has multiple layers of security.
+ It houses multiple unique modes within, allowing you to accomplish tasks with ease by simply consulting the provided help guide.

## What can you do?
+ Enhanced Security: By securely storing passwords as secrets, your sensitive credentials are now exponentially more secure.
+ Automated Passwords and Secrets: Experience the convenience of generating and storing complex passwords and secrets automatically.
+ User-Focused: Our project is tailored to provide an optimized user experience, making password management effortless and intuitive.

## How to Install?
+ You can download it from the Releases section
+ After downloading, make sure to set the location of the exe file and add the path of that configured location to your PATH file. The database will be automatically created in `C:\Users\<username>\gosecDB` PATH.
+ Note: Please be cautious while manipulating the PATH environment variable, as it directly affects how your operating system finds and executes files.
+ When downloading an executable (exe) file, you might receive a virus warning, but this can be misleading. If you'd like, you can examine the code from the source to verify.
+ If the virus threat warning doesn't go away, you can download the source code and create the exe file by writing the code below.
+ `go build main.go` After writing this, the exe file will be created inside the current folder.

## First Open
+ When you open it for the first time, gosec will ask you for a master key.
+ This is an extra security. It is useful not to forget this master key.
+ Even if you forget, you can check it in settings.json.
+  You might ask, "Why expose the master key openly in the source code?"
+ Well, imagine someone developing a program based on this source code. Let's say they are building an API and they want to decrypt their own passwords, hosting a site locally for easy access. To achieve this, they would need access to the master key. However, if the master key is embedded within the binary, it could create complications.
+ Therefore, I have incorporated this approach to address such a scenario.


## Modes
+ The application comprises five primary modes: Register, Config, Key, DeleteUser and Password.
+ All modes feature a single global option, namely -P. This argument will prompt for the password you entered during registration. It will be requested in both the password and config modes, including their sub-modes. This mechanism can be likened to an authentication check, helping us discern which user is performing the operation.
+ Now let's explain them in order

### **_Register Mode_**
+ The _register_ Mode can accept up to five arguments:
    + **-u, --user**: Input your username.
    + **-p, --pass**: Input your password.
    + **-cp, --cpass**: Confirm your password.
    + **-s, --secret**: Enter your Secret Key. Passwords added using this key will be encrypted.
    + **-gen, --generate**: Accepts one of these values: 16, 24, or 32. It will automatically create a secret for you.
+ Please note: It's imperative to remember both your password and secret key.

### **_Config Mode_**
+ The _config_ Mode can accept up to two arguments:
    + **-k, --key**: You can modify the secret you entered during registration.
    + **-req, --required**: With this, you can enable or disable additional security features in certain modes used in the password section. If you write "true", it will enable them; if you write "false", it will disable them.

### **_Key Mode_**
+ The _key_ Mode can accept up to one arguments:
    + **-gen, --generate**: Accepts one of these values: 16, 24, or 32. It will automatically create a secret or password for you.

### **_Delete User Mode_**
+ The _deleteuser_ Mode can accept up to one arguments:
    + **-p, --pass**: Deletes the user with the provided password.

### **_Password Mode_**
+ The _password_ Mode has four four distinct sub-modes: create, read, delete, and dump.
    + **_Create_**: This mode enables you to generate and store new passwords for various accounts and services. It ensures the creation of strong, unique passwords. The _create_ Mode can accept up to four arguments
        + **-t, --title**: You can add a title to specify what your password is for.
        + **-u, --url**: To indicate which website the password is associated with, you can include the URL.
        + **-p, --pas**: Enter your password.
        + **--generate**: If you don't want to create your password manually, it will generate a random 16-character password for you.
    + **_Read_**: In this mode, you can retrieve and view the passwords associated with specific accounts or services that you have stored in the system. This facilitates easy access to your stored credentials. The _read_ Mode can accept up to six arguments
        + **-i, --id**: Enables you to select the password to be used in the Read mode via its passwordId.
        + **-t, --title**: Similarly, allows you to select a password using its title.
        + **-s, --secret**: If the extra security in the config is set to true, the program might ask you to input the secret you provided during registration for certain arguments.
        + **--list**: Requires no input; you can simply call it with --list. It will list all your passwords.
        + **--open**: After entering -i or -t, if you add --open, it will open the URL associated with the password in your Chrome browser (Chrome needs to be installed).
        + **--copy**: Requires extra security. you need to provide -s and also include either -i or -t.
        + _Note_: --open and --copy can be used together `gosec password read -P <password> -i 3 --copy --open`.
    + **_Delete_**: The delete mode empowers you to remove passwords from the system for accounts or services that are no longer relevant. This helps in maintaining an organized and up-to-date password repository. The _delete_ Mode can accept up to two arguments
        + **-i, --id**: Deletes the password with the entered passwordId from the database.
        + **--all**: Requires no input; deletes all passwords.
    + **_Dump_**: The dump mode allows you to export or display all stored password information. This is useful for creating backups or managing your password data comprehensively. The _read_ Mode can accept up to tthree arguments
        + **-s, --secret**: You need to input your user secret for extra security.
        + **-p ,--path**: You can specify the path for the dump. By default, it will be saved where the binary file is located.
        + **--out**: Saves all passwords to the specified path.

# Contact

[<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/github.svg' alt='github' height='40'>](https://github.com/cetinboran)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/linkedin.svg' alt='linkedin' height='40'>](https://www.linkedin.com/in/cetinboran-mesum/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/instagram.svg' alt='instagram' height='40'>](https://www.instagram.com/2023an_m/)  [<img src='https://cdn.jsdelivr.net/npm/simple-icons@3.0.1/icons/twitter.svg' alt='twitter' height='40'>](https://twitter.com/2023anM)  





