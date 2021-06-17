var socket = new WebSocket("ws://localhost:2020/ws")

function active() {
    document.getElementById('status').innerHTML = "Online";
}

console.log("Attempting connection");

socket.onopen = () => {
    console.log("Successfully connected");
    document.getElementById('status').innerHTML = "Online";
    socket.send("Hi from client")
}
socket.onclose = (event) => {
    console.log("Socket closed: ", event)
    document.getElementById('status').innerHTML = "Offline";
}
socket.onmessage = (msg) => {
    console.log(msg);
}
socket.onerror = (error) => {
    console.log("Socket error:", error);
}


function getCookie(cname) {
    var name = cname + "=";
    var decodedCookie = decodeURIComponent(document.cookie);
    var ca = decodedCookie.split(';');
    for (var i = 0; i < ca.length; i++) {
        var c = ca[i];
        while (c.charAt(0) == ' ') {
            c = c.substring(1);
        }
        if (c.indexOf(name) == 0) {
            return c.substring(name.length, c.length);
        }
    }
    return "";
}

function help() {
    let content = document.getElementById("msg").value;
    let userId = getCookie("user_id");
    let msgContainer = document.querySelector("#msg-container");
    let date = Date.now();
    let time = new Date(date).toLocaleTimeString();
    console.log(time);
    let reps = {
        content,
        userId: getCookie("user_id"),
        time
    }
    rep = JSON.stringify(reps)
    console.log(rep);

    let template = `<div class="conversation-list">
                        <div class="chat-avatar">
                            <img src="assets/images/users/avatar-1.jpg" alt="">
                        </div>

                        <div class="user-chat-content">
                            <div class="ctext-wrap">
                                <div class="ctext-wrap-content">
                                    <p class="mb-0">
                                        ${content}
                                    </p>
                                    <p class="chat-time mb-0"><i class="ri-time-line align-middle"></i> <span class="align-middle">${time}</span></p>
                                </div>
                                <div class="dropdown align-self-start">
                                    <a class="dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-haspopup="true" aria-expanded="false">
                                        <i class="ri-more-2-fill"></i>
                                    </a>
                                    <div class="dropdown-menu">
                                        <a class="dropdown-item" href="#">Copy <i class="ri-file-copy-line float-end text-muted"></i></a>
                                        <a class="dropdown-item" href="#">Save <i class="ri-save-line float-end text-muted"></i></a>
                                        <a class="dropdown-item" href="#">Forward <i class="ri-chat-forward-line float-end text-muted"></i></a>
                                        <a class="dropdown-item" href="#">Delete <i class="ri-delete-bin-line float-end text-muted"></i></a>
                                    </div>
                                </div>
                            </div>

                            <div class="conversation-name">${userId}</div>
                        </div>

                    </div>`

    let msg = document.createElement('li');
    msg.classList.add("right");
    msg.innerHTML = template;

    msgContainer.append(msg)
    socket.send(JSON.stringify(rep))
}