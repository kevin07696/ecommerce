Use Case 1: Create User
---
The Create User use case allows a user to create a new account in the system.

Actors
---
* User

Preconditions
---
* The user is not logged in
* The user has access to the signup form

Procedure
---
1. The user navigates to the signup form.
2. The user enters email, username, and personal info into the form.
3. The system sends a one time password (otp) to the user's email.
4. The user submits the form with the otp.
5. The system validates the otp.
6. If the otp is valid, the system parses the user's inputs from the form.
7. The system validates the user's inputs
8. If the input is valid, the user is created and logged in.

Postconditions
---
* The user has a new account in the system.
* The user is logged in to the system.

Exception Paths
---
* **When the input is invalid:** the form displays inline error messages to the user.
* **When the email address is already in use:** the system displays an inline error message.
* **When the one time password is incorrect:** the form displays inline otp error message to the user.
* **When there is a problem with the system (e.g. a database error):** the system displays an error message to the user.