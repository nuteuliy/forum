document.addEventListener("DOMContentLoaded", function() {
    // Handle toggle password visibility
    const passwordInput = document.getElementById("password");
    const toggleButton = document.getElementById("toggle-password");

    if (toggleButton && passwordInput) {
        toggleButton.addEventListener("click", function() {
            if (passwordInput.type === "password") {
                passwordInput.type = "text";
                toggleButton.textContent = "Hide";
            } else {
                passwordInput.type = "password";
                toggleButton.textContent = "Show";
            }
        });
    }
    
    const postForm = document.getElementById('create-post-form')
    
    if (postForm) {
        postForm.addEventListener('submit', function(event) {
            event.preventDefault(); // Prevent default form submission

            // Collect form data
            const title = document.getElementById("title").value;
            const content = document.getElementById("content").value;
            
            // Create a user object
            const postData = {
                title: title,
                content: content
            };

            // Send the JSON data via fetch
            fetch('http://localhost:8080/create-post', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(postData) // Convert the JavaScript object to a JSON string
            })
            .then(response => {
                
                if (response.redirected) {
                    window.location.href = response.url; // Redirect to login or another page
                } else {
                    return response.json(); // Assuming error messages are returned as text
                }
            })
            .then(data => {
                // Update your form with error messages from the server
                
                console.log(data)
                const wordCountError = document.querySelector('.wordCountError');
                

                // Display errors
                wordCountError.textContent = data.countError || '';
                wordCountError.style.display = 'block'
            })
            .catch(error => {
                console.error('Error during registration:', error);
            });
        });
    }
    // Handle registration form
    const registrationForm = document.getElementById('registrationForm');
    if (registrationForm) {
        registrationForm.addEventListener('submit', function(event) {
            event.preventDefault(); // Prevent default form submission

            // Collect form data
            const email = document.getElementById('email').value;
            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            // Create a user object
            const user = {
                email: email,
                username: username,
                password: password
            };

            // Send the JSON data via fetch
            fetch('http://localhost:8080/register', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(user) // Convert the JavaScript object to a JSON string
            })
            .then(response => {
                if (response.redirected) {
                    window.location.href = response.url; // Redirect to login or another page
                } else {
                    return response.json(); // Assuming error messages are returned as text
                }
            })
            .then(data => {
                // Update your form with error messages from the server
                const emailError = document.querySelector('.email-error');
                const usernameError = document.querySelector('.username-error');
                const passwordError = document.querySelector('.password-error');

                // Display errors
                emailError.textContent = data.emailError || '';
                usernameError.textContent = data.usernameError || '';
                passwordError.textContent = data.passwordError || '';
            })
            .catch(error => {
                console.error('Error during registration:', error);
            });
        });
    }

    // Handle login form
    const loginForm = document.getElementById('loginForm');
    if (loginForm) {
        loginForm.addEventListener('submit', function(event) {
            event.preventDefault(); // Prevent default form submission

            // Collect login form data
            const username = document.getElementById('username').value;
            const password = document.getElementById('password').value;

            // Create a login object
            const userLog = {
                username: username,
                password: password
            };

            // Send the JSON data via fetch for login
            fetch('http://localhost:8080/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(userLog)
            })
            .then(response => {
                if (response.redirected) {
                    window.location.href = response.url; // Redirect to another page
                } else {
                    return response.json(); // Assuming error messages are returned as text
                }
                console.log(response)
            })
            .then(data => {
                
                // Update your form with error messages from the server
                const usernameError = document.querySelector('.username-error-login');
                const passwordError = document.querySelector('.password-error-login');

                // Display errors
                usernameError.textContent = data.usernameError || '';
                passwordError.textContent = data.passwordError || '';
            })
            .catch(error => {
                console.error('Error during login:', error);
            });
        });
    }
});
