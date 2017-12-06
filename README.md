# ssf
A command line tool written in Go to list and delete shared files in slack.

# Install 
go get github.com/xingye/ssf

# Usage
To use this tool to list or delete your slack shared files, you need to generate a token first.
Please refer the [slack file method page](https://api.slack.com/custom-integrations/legacy-tokens).
If you are workplace owner and admin, you can delete any file you want. If you are a member and guest,
you can only delete your own files. In order to list your own files, you can provide your user id, but 
token is required. For your convenience, you can **export** your token and user id as **ssf_token** and **ssf_user**
before using this tool.

![command](/src/command.png)

# Todo
1. Make progress bar in terminal when listing or deleting files.
2. Provide option to delete one file or specific files meeting condition(such as older than 7 days).
