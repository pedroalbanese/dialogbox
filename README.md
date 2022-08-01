# dialogbox
## Cross-platform Dialog Boxes Tool written in Go 
Custom Input/Dialog boxes to get user input/choice for bash/batch scripting automation and interaction. A graphical control element in the form of a small window that communicates information to the user and prompts them for a response. 

### Usage of dialogbox:
<pre>  -date
        Date selection dialog box
  -error
        Error dialog box
  -file
        File selection dialog box
  -folder
        Folder selection dialog box
  -info
        Info dialog box
  -input
        Text input box
  -pass
        Password input box
  -quest
        Question dialog box (Check the Exit code)
  -sub string
        Box subtitle (default "Subtitle")
  -title string
        Box title (default "Title")</pre>

### Examples:
Password insertion:
<pre>password=$(./dialogbox -pass -title "Password" -sub "Insert Password:")</pre>
Link insertion:
<pre>link=$(./dialogbox -input -title "Link" -sub "Insert the link:")</pre>
Question:
<pre>./dialogbox -quest -title "Verification" -sub "Are you sure?"
echo $?</pre>

Link insertion (Windows):
<pre>FOR /F %%F in ('dialogbox -input -title "Link" -sub "Insert the link:"') do (set link=%%F)</pre> 
Question (Windows):
<pre>dialogbox -quest -title "Verification" -sub "Are you sure?"
echo %ERRORLEVEL%</pre>

### Linux:
<pre>$ apt-get install zenity</pre>

## License
This project is licensed under the ISC License.
##### Copyright (c) 2020-2022 Pedro F. Albanese - ALBANESE Research Lab.
