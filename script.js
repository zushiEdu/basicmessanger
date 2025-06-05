sendButtom = document.getElementById("send")
sendForm = document.getElementById("text")

chatHistory = document.getElementById("chat")

idFrom = document.getElementById("idFrom")
idTo = document.getElementById("idTo")

const url = "http://10.0.0.243:2345/message/"

function sendMessage() {
    var options = {
        fromUser: idFrom.value,
        toUser: idTo.value,
        message: sendForm.value
    };

    fetch(url, {
        method: 'POST',
        body: JSON.stringify(options)
    })
        .then(response => response.json())
        .then(data => window.alert("Message Sent."))
        .catch(error => console.error('Error:', error));
}

function getMessages() {
    var options = new URLSearchParams({ toUser: idFrom.value })

    var responseData;

    fetch(url + `?${options}`, {
        method: 'GET',
    })
        .then(response => response.json())
        .then(data => {
            chatHistory.innerHTML = '';

            if (data.data != null) {
                data.data.forEach((message, index) => {
                    var newMessage = document.createElement('p')
                    newMessage.textContent = `${index + 1} - ${message}`
                    chatHistory.appendChild(newMessage)
                })
            };
        })
        .catch(error => console.error('Error:', error));
}

setInterval(getMessages, 100)