const url = "http://10.0.0.243:1234/"

let loginPanel = document.getElementById('loginAccount')
let createAccountPanel = document.getElementById('createAccount')
let chatSelectionPanel = document.getElementById('chatSelection')
let messageForm = document.getElementById('messageForm')
let message = document.getElementById('text')
let chatBox = document.getElementById('chat')
let contactNew = document.getElementById('contactNew')

let menuState = "loginScreen"

setMenuState();
function setMenuState() {
    loginPanel.style.display = "none"
    messageForm.style.display = "none"
    createAccountPanel.style.display = "none"
    chatSelectionPanel.style.display = "none"
    contactNew.style.display = "none"

    if (menuState === "chatSelection") {
        chatSelectionPanel.style.display = "grid"
        refreshChatSelection()
    } else if (menuState === "chatWindow") {
        messageForm.style.display = "grid"
        refreshChatWindow()
    } else if (menuState === "createAccountWindow") {
        createAccountPanel.style.display = "grid"
    } else if (menuState === "loginScreen") {
        loginPanel.style.display = "grid"
    } else if (menuState === "contactNew") {
        contactNew.style.display = "grid"
    }
}

function triggerChatSelection() {
    menuState = "chatSelection"
}

let searchByEmail = document.getElementById('searchByEmail')
function triggerContactNew() {
    menuState = "contactNew"
    searchByEmail.value = ""
}

// Switches between create account and login screens
function toggleLogin() {
    if (menuState === "createAccountWindow") {
        menuState = "loginScreen"
    } else if (menuState === "loginScreen") {
        menuState = "createAccountWindow"
    }
}

let emailLoginForm = document.getElementById('emailLogin');
let passwordLoginForm = document.getElementById('passwordLogin');

// Gets called when login button is pressed
let token
function tryLogin() {
    console.log("Trying to log in...")

    let options = {
        email: emailLoginForm.value,
        password: passwordLoginForm.value
    };

    fetch(url + "login", {
        method: 'POST',
        body: JSON.stringify(options)
    })
        .then(response => {
            if (response.status === 401) {
                alert("Incorrect email or password")
                throw new Error("Unauthorized")
            }
            if (!response.ok) {
                throw new Error("Could not log in")
            }

            return response.json();
        })
        .then(data => {
            localStorage.setItem("token", data.token);
            emailLoginForm.value = ""
            passwordLoginForm.value = ""
            menuState = "chatSelection"
        })
        .catch(error => console.error('Error:', error));
}

let firstNameCreate = document.getElementById('firstNameCreate')
let lastNameCreate = document.getElementById('lastNameCreate')
let emailCreate = document.getElementById('emailCreate')
let passwordCreate = document.getElementById('passwordCreate')
let confirmPasswordCreate = document.getElementById('passwordCreateConfirm')
function tryCreateAccount() {
    let options = { firstName: firstNameCreate.value, lastName: lastNameCreate.value, email: emailCreate.value, password: passwordCreate.value };

    console.log(JSON.stringify(options))

    if (passwordCreate.value === confirmPasswordCreate.value) {
        fetch(url + `users`, {
            method: 'POST',
            body: JSON.stringify(options),
        })
            .then(response => {
                if (response.status === 409) {
                    alert("Account with email already exists")
                    throw new Error("Could not create account")
                } else if (response.ok) {
                    menuState = "loginScreen"
                    alert("Created new account successfully.")

                    firstNameCreate.value = ""
                    lastNameCreate.value = ""
                    emailCreate.value = ""
                    passwordCreate.value = ""
                    confirmPasswordCreate = ""
                }
            })
            .catch(error => console.error('Error:', error));
    } else {
        alert("Passwords do not match")
    }
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
                    let name = document.createElement("a")
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
    menuState = "chatWindow"
    globalID = id
    let personName = document.getElementById('messageBarName')
    personName.innerText = name
    refreshChatWindow()
}

function openUndefinedChat(email) {

    let options = new URLSearchParams({ mode: "single", email: searchByEmail.value })

    fetch(url + `users?${options}`, {
        method: 'GET',
        headers: {
            'Content-Type': 'application/json',
            'Authorization': `Bearer ${localStorage.getItem('token')}`
        },
    })
        .then(response => response.json())
        .then(data => {
            if (data.message === "User not found") {
                alert("User not found")
            } else {
                openDefinedChat(data.data[0].id, data.data[0].firstName + " " + data.data[0].lastName)
                menuState = "chatWindow"
            }
        })
        .catch(error => console.error('Error:', error));

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
        .then(data => {
            message.value = ""
            refreshChatWindow()
        })
        .catch(error => console.error('Error:', error));

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

setInterval(setMenuState, 500)