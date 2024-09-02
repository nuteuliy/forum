document.addEventListener("DOMContentLoaded", function() {
    const passwordInput = document.getElementById("password");
    const toggleButton = document.getElementById("toggle-password");

    toggleButton.addEventListener("click", function() {
        if (passwordInput.type === "password") {
           
            passwordInput.type = "text";
            toggleButton.textContent = "Hide";
        } else {
            passwordInput.type = "password";
            toggleButton.textContent = "Show";
        }
    });
    document.getElementById('registrationForm').addEventListener('submit', function(event) {
      
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
            console.log(data);
            // Here, update your form with error messages from the server
                const emailError = document.querySelector('.email-error');
                const usernameError = document.querySelector('.username-error');
                const passwordError = document.querySelector('.password-error');
                console.log(emailError,data.emailError)

                // Display errors
                emailError.textContent = data.emailError|| '';
                usernameError.textContent = data.usernameError || '';
                passwordError.textContent = data.passwordError || '';
        })
        .catch(error => {
            console.error('Error:', error);
        });
    });
});