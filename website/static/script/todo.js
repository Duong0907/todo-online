


// ---------------
//      Logout
// ---------------
function logout() {
    // Set expiring time to a past time to delete the cookie
    document.cookie = "todo_api_jwt=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;";
    location.reload();
}



// ---------------
//      Login
// ---------------
var loginModal = document.querySelector('#login-container #id01');
var usernameLoginInput = document.querySelector('#login-container .uname-input');
var passwordLoginInput = document.querySelector('#login-container .psw-input');
var usernameSignupInput = document.querySelector('#signup-container .uname-input');
var emailSignupInput = document.querySelector('#signup-container .email-input');
var passwordSignupInput = document.querySelector('#signup-container .psw-input');


window.addEventListener('click', (event) => {
    if (event.target == loginModal) {
        loginModal.style.display = "none";
    }
})

function login() {
    credentials = {
        'username': usernameLoginInput.value,
        'password': passwordLoginInput.value
    }

    loginInDB(credentials)
        .then(res => {
            if (!res['error']) {
                loginModal.style.display = 'none';
                startWebsite();
                location.reload();
                clearModal(document.getElementById('id01'));
            } else {
                alert(res['message']);
            }
        })
}

function loginInDB(credentials) {
    var username = credentials['username'];
    var password = credentials['password'];
    return fetch(rootApi + `users/login?username=${username}&password=${password}`, { method: 'POST' })
        .then(res => res.json())
}

function clearModal(modal) {
    modal.style.display = 'none';

    var items = modal.getElementsByTagName('input');
    for (let i = 0; i < items.length; i++) {
        items[i].value = '';
    }
}

function listCookies() {
    var theCookies = document.cookie.split(';');
    var aString = '';
    for (var i = 1; i <= theCookies.length; i++) {
        aString += i + ' ' + theCookies[i - 1] + "\n";
    }
    return aString;
}

function getCookie(name) {
    const cDecoded = decodeURIComponent(document.cookie);
    const cArray = cDecoded.split('; ');
    let jwtString = null;

    cArray.forEach(element => {
        if (element.indexOf(name) == 0) {
            jwtString = element.substring(name.length + 1);
        }
    })
    return jwtString;
}




// -----------------
//      Signup
// -----------------

var signupModal = document.querySelector('#signup-container #id02');

window.addEventListener('click', (event) => {
    if (event.target == signupModal) {
        signupModal.style.display = "none";
    }
})

function createUserInDB(user) {
    var options = {
        method: 'POST'
    }
    var endpoint = `users/?username=${user.username}&email=${user.email}&password=${user.password}`
    return fetch(rootApi + endpoint, options)
        .then(res => res.json())
}

function signup() {
    var user = {
        username: usernameSignupInput.value,
        email: emailSignupInput.value,
        password: passwordSignupInput.value,
    }
    console.log(user)
    createUserInDB(user)
        .then(res => {
            if (res['error']) {
                alert(res['message']);
            } else {
                clearModal(document.getElementById('id02'));
            }
        })
}

// ---------------------------
//             Alert 
// ---------------------------

function alert(message) {
    var alertElement = document.getElementsByClassName('alert')[0];
    console.log(alertElement)
    document.querySelector('.alert .message').innerHTML = message;
    alertElement.style.display = 'block';
    setTimeout(() => {
        alertElement.style.display = 'none';
    }, 1500);
}

// ----------------------------
//              Todo
// ----------------------------

// Todo container
const todoContainer = document.getElementById('todo-container');
const myInput = document.querySelector('#todo-container #myInput');
const myUL = document.querySelector('#todo-container #myUL');


var rootApi = `${window.location}/api/`

// Interacting data in database functions
function getDataFromDB(endpoint) {
    let options = {
        headers: {
            'Authorization': 'Bearer ' + getCookie('todo_api_jwt')
        }
    }
    return fetch(rootApi + endpoint, options)
        .then(response => response.json())
}

function addTodoToDB(todo) {
    let options = {
        method: 'POST',
        headers: {
            'Authorization': 'Bearer ' + getCookie('todo_api_jwt')
        },
        body: JSON.stringify(todo)
    }
    return fetch(rootApi + 'todos/', options)
        .then(res => res.json())
}

function removeTodoFromDB(id) {
    var options = {
        method: 'DELETE',
        headers: {
            'Authorization': 'Bearer ' + getCookie('todo_api_jwt')
        }
    }
    return fetch(rootApi + 'todos/?id=' + id, options)
        .then(res => res.json())
}

function toggleCheckTodoToDB(id) {
    var options = {
        method: 'PUT',
        headers: {
            'Authorization': 'Bearer ' + getCookie('todo_api_jwt')
        }
    }
    return fetch(rootApi + 'todos/?id=' + id, options)
        .then(res => res.json())
}


// Adding event functions
function addCheckedState(taskItem) {
    taskItem.addEventListener('click', (event) => {
        event.target.classList.toggle('checked');
        toggleCheckTodoToDB(event.target.getAttribute('id'))
    });
}

function addCloseButton(taskItem) {
    let closeButton = document.createElement('div');
    closeButton.innerHTML = 'X';
    closeButton.classList.add('close');
    closeButton.setAttribute("style", "font-family: cursive; font-size: 15px;");
    closeButton.setAttribute('id', taskItem.getAttribute('id'))

    // click event
    closeButton.addEventListener('click', (event) => {
        event.target.parentElement.remove();
        removeTodoFromDB(event.target.getAttribute('id'))
    })

    taskItem.appendChild(closeButton);
}


function setEnterSubmitInput() {
    myInput.addEventListener('keydown', event => {
        if (event.key === 'Enter') {
            newElement()
        }
    })
}


function newElement() {
    let taskContent = myInput.value;
    if (taskContent == "") {
        return
    }

    myInput.value = ""
    let newTaskElement = document.createElement('li');
    newTaskElement.innerHTML = taskContent;


    var todo = {
        "title": taskContent,
        "status": 'Todo'
    }
    addTodoToDB(todo)
        .then(res => {
            console.log(res)
            if (!res['error']) {
                newTaskElement.setAttribute('id', res['data']['todo']['id'])
            }
            addCheckedState(newTaskElement);
            addCloseButton(newTaskElement);
            myUL.appendChild(newTaskElement);
        })
}


function renderAnItem(item) {
    let newTaskElement = document.createElement('li');
    newTaskElement.innerHTML = item["title"];
    if (item["status"] === 'Done')
        newTaskElement.classList.add("checked");
    newTaskElement.setAttribute("id", item["id"]);

    addCheckedState(newTaskElement);
    addCloseButton(newTaskElement);

    myUL.appendChild(newTaskElement);
}


// Starting the website
function startWebsite() {
    getDataFromDB('users/')
        .then(res => {
            if (!res['error']) {
                // Display user icon
                document.getElementById('login-container').style.display = 'none';
                document.getElementById('signup-container').style.display = 'none';
                document.getElementById('account-container').style.display = 'inline-block';
                document.getElementById('logout-container').style.display = 'inline-block';
                document.getElementById('account-name').innerHTML = res['data']['user']['username'];

                // Show todo data of the user
                getDataFromDB('todos/')
                    .then(response => {
                        response['data']['todos'].forEach(item => {
                            renderAnItem(item);
                        });
                        setEnterSubmitInput();
                    })
            }
        })
}
startWebsite()