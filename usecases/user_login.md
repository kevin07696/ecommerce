Use Case 1: Login User
---
The Login User allows a user to log in to the system without a password.

Actors
---
* User

Preconditions
---
* The user is not logged in
* The user has access to the login form
* The user has access to his/her email

Procedure
---
1. The user navigates to the login form.
2. The user enters email into the form.
3. The system sends a one time password (otp) to the user's email.
4. The user submits the otp.
5. The system validates the otp.
6. If the otp is valid, the system parses the user's inputs from the form.
7. The system validates the user's inputs
8. If the input is valid, the system finds the user data using the unique email address.
9. Navigate to the home page for the user.

Postconditions
---
* The user is logged in to the system.

Exception Paths
---
* **When the email is invalid:** the form displays inline error messages to the user.
* **When the one time password is incorrect:** the form displays inline otp error message to the user.
* **When the email is not found:** the form displays inline error messages to the user.
* **When there is a problem with the system (e.g. a database error):** the system displays an error message to the user.