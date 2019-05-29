# Finance Chat

### How to run
In order to run the web chat you need to run `go run main.go` on src of both proyects: _finance-chat_ and _finance-chat-bot_. You will need the ports :8000 and :8001 free. Also, you may need to install dependencies with npm. Open the public folder on _finance-chat_ and run `npm install`.

#### Requirements
  - github.com/gin-gonic/gin
  - github.com/gorilla/websocket
  - A mysql running on localhost with user and password root

#### Unfinished tasks
  - I didn't do any bonuses
  - The bot is not finished. It's almost done but I had a problem with WSS communication. When you try to send a message to it, it finish the execution with an error of protocols
  - I didn't test anything. I think that a feature is not finished until you test it, so this point makes me really sad. I ran out of time and i couldn't test anything.

#### Other comments

I think that the project could be much better done than it is, but the time played against me. I thought that this weekend i could have a lot of free time, but university made that not happen. These are not excuses, only clarifications that I think this has many opportunities for improvement, like the structure of the project, the repeated code, the validations, the configurations (I have a password in a const file) and the versioned by git.

Beyond this, I await your feedback and any questions I'm available.