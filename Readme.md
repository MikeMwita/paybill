# PayBill

PayBill is a comprehensive solution crafted to streamline the intricacies of bill payment, tracking, and approval seamlessly. It offers finance teams a centralized dashboard, enhancing their operational efficiency. Key features include bulk upload, smart scanning, and automated bill payments. This service supports both inter-currency and international payments, seamlessly synchronizing all transactions with Xero.

## Technologies

This project is built with the following technologies:

- Golang
- PostgreSQL
- Redis

## Features

- **User authentication and authorization**: This feature allows users to register, login, logout, and refresh their access tokens. It also allows users to reset their passwords, verify their emails, and update their profiles. This feature uses JWT or OAuth 2.0 to secure your endpoints and manage different user roles and permissions.
- **Bill management**: This feature allows users to create, update, delete, and pay their bills. It also allows users to upload bills in bulk, scan bills using OCR, and automate bill payments. This feature supports both inter-currency and international payments, seamlessly synchronizing all transactions .
- **Notification system**: This feature allows users to send and receive email or SMS alerts when their bills are due, paid, or overdue. It also allows users to delete their notifications. 
- **Integration system**: This feature allows users to connect and disconnect with other accounting software or platforms like QuickBooks, FreshBooks, or Stripe. It also allows users to get all their integrations and a single integration. This feature enables users to support more payment methods and currencies, as well as sync their data with other systems.
- **Reporting and analytics system**: This feature allows users to generate, view, and delete reports on their spending habits, payment history, and cash flow. It also allows users to get a single report. This feature uses libraries like Gonum or Plotly to generate insights and visualizations.
- **Feedback and rating system**: This feature allows users to rate and review the merchants they pay. It also allows users to get and delete their feedback. This feature helps users to improve their service quality and customer satisfaction.
- **Goal savings and savings account**: This feature allows users to create, update, delete, and transfer money to or from their savings goals. It also allows users to withdraw money from their savings goals. This feature helps users to save money for a specific purpose, such as an emergency fund, a vacation, or a down payment.


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





| Endpoint | Method | Description |
| --- | --- | --- |
| /api/goals | POST | Create a new savings goal |
| /api/goals | GET | Get all savings goals |
| /api/goals/:id | GET | Get a single savings goal |
| /api/goals/:id | PUT | Update a savings goal |
| /api/goals/:id | DELETE | Delete a savings goal |
| /api/goals/:id/transfer | POST | Transfer money to or from a savings goal |
| /api/goals/:id/withdraw | POST | Withdraw money from a savings goal |





## License

This project is licensed under the MIT License. See the [LICENSE] file for more details.

