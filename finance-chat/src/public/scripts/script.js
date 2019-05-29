import axios from '../node_modules/@bundled-es-modules/axios/axios.js';

new Vue({
    el: '#finance-chat',

    data: {
        ws: null, // Our websocket
        newMsg: '', // Holds new messages to be sent to the server
        chatContent: '', // A running list of chat messages displayed on the screen
        messages: [],
        user: {
            username: null, // Our username
            password: null
        },
        joined: false // True if email and username have been filled in
    },
    created: function () {
        var self = this;
        this.ws = new WebSocket('wss://' + window.location.host + '/wss');
        this.ws.addEventListener('message', function (e) {
            var msg = JSON.parse(e.data);
            self.addMessageToMessages(msg);
            self.parseMessagesToChat();

            var element = document.getElementById('chat-messages');
            element.scrollTop = element.scrollHeight; // Auto scroll to the bottom
        });
    },
    methods: {
        send: function () {
            if (this.newMsg !== '') {
                let date = new Date();
                this.ws.send(
                    JSON.stringify({
                            username: this.user.username,
                            message: $('<p>').html(this.newMsg).text(), // Strip out html
                            time: date.getHours() + ':' + date.getMinutes() + ':' + date.getSeconds()
                        }
                    ));
                this.newMsg = ''; // Reset newMsg
            }
        },
        signin: function () {
            if (!this.user.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            if (!this.user.password) {
                Materialize.toast('You must choose a password', 2000);
                return
            }
           axios({
               method: 'put',
               url: "https://localhost:8000/api/users/login",
               data: {
                   username: this.user.username,
                   password: this.user.password // as is an http server, it goes encrypted
               }
           })
               .then(() => {
                   this.username = $('<p>').html(this.username).text();
                   this.joined = true;
               })
               .catch(err => {
                   Materialize.toast(`An error has occurred when you tried to log in: ${err.response.data.error}`, 5000)
               });
        },
        signup: function() {
            if (!this.user.username) {
                Materialize.toast('You must choose a username', 2000);
                return
            }
            if (!this.user.password) {
                Materialize.toast('You must choose a password', 2000);
                return
            }
            axios({
                method: 'post',
                url: "https://localhost:8000/api/users",
                data: {
                    username: this.user.username,
                    password: this.user.password // as is an http server, it goes encrypted
                }
            })
                .then(() => {
                    this.username = $('<p>').html(this.username).text();
                    this.joined = true;
                })
                .catch(err => {
                    Materialize.toast(`An error has occurred when you tried to log in: ${err.response.data.error}`, 5000)
                });
        },
        addMessageToMessages: function(message) {
            this.messages.push(message);
            this.messages.sort(function(a, b) {
                return a.time - b.time;
            });
            if (this.messages.length > 5) {
                this.messages.shift();
            }
        },
        parseMessagesToChat: function () {
            this.chatContent = ' ';
            var self = this;
            this.messages.forEach(function(message) {
                self.chatContent += '<div class="chip">'
                    + message.username
                    + '</div>'
                    + message.message + '<div class="time">' + message.time + '</div><br/>';
            });
        }
    }
});
