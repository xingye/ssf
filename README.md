# ssf
A command line tool written in Go to list and delete shared files in slack.

# Install 
1. clone this repo
2. cd ssf/
3. go get ./...
4. go install

**Note** add your **go-workplace/bin** directory to path, so that you can exec **ssf** command in everywhere.

# Purpose
Since [message and storage limits on the free plan](https://get.slack.help/hc/en-us/articles/115002422943-Message-and-storage-limits-on-the-Free-plan) of slack, workspaces on the free plan share a total of 5GB of file storage space. 
When a workplace file storage space exceeds the limit, you can not upload file any more. You need to go to [website](https://get.slack.help/hc/en-us/articles/218159688-Delete-shared-files) and delete files one by one. It is very inconveniences. So I write this tool to bulk delete files and make life more easy.

# Usage
To use this tool to list or delete your slack shared files, you need to generate a token first.
Please refer the [slack file method page](https://api.slack.com/custom-integrations/legacy-tokens).
If you are workplace owner and admin, you can delete any file you want. If you are a member and guest,
you can only delete your own files. In order to list your own files, you can provide your user id, but 
token is required. If u option missing, it will list all the flies of the workplace. For your convenience, 
you can **export** your token as **ssf_token** before using this tool. Btw, You can use command**ssf echo** to 
show your slack user id.

![command](/src/command1.png)
![command](/src/command2.png)

# Todo
1. Provide option to delete one file or specific files meeting condition(such as older than 7 days).
