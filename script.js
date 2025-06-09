const url = "http://127.0.0.1:1234/"

let loginPanel = document.getElementById('loginAccount')
let createAccountPanel = document.getElementById('createAccount')
let chatSelectionPanel = document.getElementById('chatSelection')
let messageForm = document.getElementById('messageForm')
let message = document.getElementById('text')
let chatBox = document.getElementById('chat')

let loginDisplay = true;

if (localStorage.getItem("logged") === null) {
    console.log("First time use")
    localStorage.setItem("logged", false)
} else {
    console.log("Logged in before with state:", localStorage.getItem("logged"))
    if (localStorage.getItem("logged") == "true") {
        loginPanel.style.display = "none"
        chatSelectionPanel.style.display = "grid"
    } else {
        loginPanel.style.display = "grid"
        chatSelectionPanel.style.display = "none"
    }
}

function toggleLogin() {
    if (loginDisplay) {
        loginPanel.style.display = "none";
        createAccountPanel.style.display = "grid";
        loginDisplay = false;
    } else {
        loginPanel.style.display = "grid";
        createAccountPanel.style.display = "none";
        loginDisplay = true;
    }
}

let emailLoginForm = document.getElementById('emailLogin');
let passwordLoginForm = document.getElementById('passwordLogin');

let token
function tryLogin() {
    let options = {
        email: emailLoginForm.value,
        password: passwordLoginForm.value
    };

    fetch(url + "login", {
        method: 'POST',
        body: JSON.stringify(options)
    })
        .then(response => response.json())
        .then(data => {
            localStorage.setItem("logged", true);
            localStorage.setItem("token", data.token);
        })
        .catch(error => console.error('Error:', error));
}

function refreshChatSelection() {
    let options = new URLSearchParams({ mode: "multi" })

    fetch(url + `users?${options}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
    })
        .then(response => response.json())
        .then(data => {
            if (Array.isArray(data?.data) && data.data.length > 0) {
                let elementsToRemove = chatSelectionPanel.querySelectorAll('[id="chatOption"]');
                elementsToRemove.forEach(child => { child.remove() })

                for (let i = 0; i < (data.data.length); i++) {
                    let chatOption = document.createElement("div")
                    chatOption.id = "chatOption"
                    let name = document.createElement("p")
                    let nameText = data.data[i].firstName + " " + data.data[i].lastName
                    name.innerText = nameText
                    chatOption.onclick = function () {
                        openDefinedChat(data.data[i].id, nameText)
                    }
                    chatOption.appendChild(name)
                    chatSelectionPanel.appendChild(chatOption)
                }
            }
        })
        .catch(error => console.error('Error:', error));

}

let globalID;
function openDefinedChat(id, name) {
    globalID = id
    chatSelectionPanel.style.display = "none"
    messageForm.style.display = "grid"
    let personName = document.getElementById('messageBarName')
    personName.innerText = name
    refreshChatWindow()
}

function sendMessage() {
    let options = {
        toUser: globalID,
        message: message.value
    };

    fetch(url + "message", {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
        body: JSON.stringify(options),
    })
        .then(response => { response.json() })
        .catch(error => console.error('Error:', error));

    refreshChatWindow
}

function refreshChatWindow() {
    let options = new URLSearchParams({ involvingUser: globalID })

    fetch(url + `message?${options}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
    })
        .then(response => response.json())
        .then(data => {
            if (Array.isArray(data?.data) && data.data.length > 0) {
                chatBox.innerHTML = ""

                for (let i = 0; i < data.data.length; i++) {
                    let bubble = document.createElement('p')
                    bubble.innerText = data.data[i].message
                    if (data.data[i].userFrom == globalID) {
                        // came from other person, put on left
                        bubble.id = "leftBubble"
                    } else {
                        bubble.id = "rightBubble"
                    }
                    chatBox.appendChild(bubble)
                }
            }
        })
        .catch(error => console.error('Error:', error));

}

refreshChatSelection()
setInterval(refreshChatSelection, 1000)
setInterval(refreshChatWindow, 1000)