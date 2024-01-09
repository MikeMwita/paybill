# PayBill

PayBill is a comprehensive solution crafted to streamline the intricacies of bill payment, tracking, and approval seamlessly. It offers finance teams a centralized dashboard, enhancing their operational efficiency. Key features include bulk upload, smart scanning, and automated bill payments. This service supports both inter-currency and international payments, seamlessly synchronizing all transactions with Xero.

## Technologies

This project is built with the following technologies:

- Golang
- MySQL
- Redis

## Installation

To install and run this project, you need to have Golang, MySQL, and Redis installed on your machine. Then follow these steps:

1. Clone this repository to your local machine: `git clone https://github.com/MikeMwita/PayBill.git`
2. Navigate to the project directory: `cd PayBill`
3. Install the dependencies: `go mod download`
4. Create a `.env` file and add your environment variables (such as database credentials, API keys, etc.)
5. Run the project: `go run main.go`

## Endpoints



| Endpoint | Method | Description |
| --- | --- | --- |
| /api/auth/register | POST | Register a new user |
| /api/auth/login | POST | Login an existing user |
| /api/auth/logout | POST | Logout the current user |
| /api/auth/refresh | POST | Refresh the access token |



| Endpoint | Method | Description |
| --- | --- | --- |
| /api/auth/reset | POST | Request a password reset link |
| /api/auth/reset/:token | PUT | Reset the password using the token |
| /api/auth/verify | POST | Request an email verification link |
| /api/auth/verify/:token | GET | Verify the email using the token |
| /api/auth/profile | GET | Get the user profile |
| /api/auth/profile | PUT | Update the user profile |



| Endpoint | Method | Description |
| --- | --- | --- |
| /api/bills | POST | Create a bill |
| /api/bills | GET | Get all bills |
| /api/bills/:id | GET | Get a single bill |
| /api/bills/:id | PUT | Update a bill |
| /api/bills/:id | DELETE | Delete a bill |
| /api/bills/:id/pay | POST | Pay a bill |
| /api/notifications | POST | Send a notification |
| /api/notifications | GET | Get all notifications |
| /api/notifications/:id | GET | Get a single notification |
| /api/notifications/:id | DELETE | Delete a notification |
| /api/integrations | POST | Connect with an accounting platform |
| /api/integrations | GET | Get all integrations |
| /api/integrations/:id | GET | Get a single integration |
| /api/integrations/:id | DELETE | Disconnect from an accounting platform |
| /api/reports | POST | Generate a report |
| /api/reports | GET | Get all reports |
| /api/reports/:id | GET | Get a single report |
| /api/reports/:id | DELETE | Delete a report |
| /api/feedback | POST | Rate a merchant |
| /api/feedback | GET | Get all feedback |
| /api/feedback/:id | GET | Get a single feedback |
| /api/feedback/:id | DELETE | Delete a feedback |


## Contributing

This project is open for contributions. If you want to contribute, please follow these steps:

1. Fork this repository and create a new branch: `git checkout -b feature-name`
2. Make your changes and commit them: `git commit -m "Add some feature"`
3. Push your branch to your forked repository: `git push origin feature-name`
4. Create a pull request and wait for review.

## License

This project is licensed under the MIT License. See the [LICENSE] file for more details.

