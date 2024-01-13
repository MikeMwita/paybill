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


# Documentation for API Endpoints

## Authentication

This section covers the endpoints related to user authentication and profile management.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| AuthApi | Register | POST /api/auth/register | Register a new user |
| AuthApi | Login | POST /api/auth/login | Login an existing user |
| AuthApi | Logout | POST /api/auth/logout | Logout the current user |
| AuthApi | Refresh | POST /api/auth/refresh | Refresh the access token |
| AuthApi | Reset | POST /api/auth/reset | Request a password reset link |
| AuthApi | ResetPassword | PUT /api/auth/reset/:token | Reset the password using the token |
| AuthApi | Verify | POST /api/auth/verify | Request an email verification link |
| AuthApi | VerifyEmail | GET /api/auth/verify/:token | Verify the email using the token |
| AuthApi | GetProfile | GET /api/auth/profile | Get the user profile |
| AuthApi | UpdateProfile | PUT /api/auth/profile | Update the user profile |

## Bills

This section covers the endpoints related to bill creation, management, and payment.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| BillApi | CreateBill | POST /api/bills | Create a bill |
| BillApi | GetBills | GET /api/bills | Get all bills |
| BillApi | GetBill | GET /api/bills/:id | Get a single bill |
| BillApi | UpdateBill | PUT /api/bills/:id | Update a bill |
| BillApi | DeleteBill | DELETE /api/bills/:id | Delete a bill |
| BillApi | PayBill | POST /api/bills/:id/pay | Pay a bill |

## Notifications

This section covers the endpoints related to notification sending and management.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| NotificationApi | SendNotification | POST /api/notifications | Send a notification |
| NotificationApi | GetNotifications | GET /api/notifications | Get all notifications |
| NotificationApi | GetNotification | GET /api/notifications/:id | Get a single notification |
| NotificationApi | DeleteNotification | DELETE /api/notifications/:id | Delete a notification |

## Integrations

This section covers the endpoints related to integration with accounting platforms.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| IntegrationApi | Connect | POST /api/integrations | Connect with an accounting platform |
| IntegrationApi | GetIntegrations | GET /api/integrations | Get all integrations |
| IntegrationApi | GetIntegration | GET /api/integrations/:id | Get a single integration |
| IntegrationApi | Disconnect | DELETE /api/integrations/:id | Disconnect from an accounting platform |

## Reports

This section covers the endpoints related to report generation and management.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| ReportApi | GenerateReport | POST /api/reports | Generate a report |
| ReportApi | GetReports | GET /api/reports | Get all reports |
| ReportApi | GetReport | GET /api/reports/:id | Get a single report |
| ReportApi | DeleteReport | DELETE /api/reports/:id | Delete a report |

## Feedback

This section covers the endpoints related to feedback rating and management.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| FeedbackApi | Rate | POST /api/feedback | Rate a merchant |
| FeedbackApi | GetFeedback | GET /api/feedback | Get all feedback |
| FeedbackApi | GetFeedback | GET /api/feedback/:id | Get a single feedback |
| FeedbackApi | DeleteFeedback | DELETE /api/feedback/:id | Delete a feedback |

## Goals

This section covers the endpoints related to savings goals creation, management, and transfer.

| Class | Method | HTTP request | Description |
| --- | --- | --- | --- |
| GoalApi | CreateGoal | POST /api/goals | Create a new savings goal |
| GoalApi | GetGoals | GET /api/goals | Get all savings goals |
| GoalApi | GetGoal | GET /api/goals/:id | Get a single savings goal |
| GoalApi | UpdateGoal | PUT /api/goals/:id | Update a savings goal |
| GoalApi | DeleteGoal | DELETE /api/goals/:id | Delete a savings goal |
| GoalApi | Transfer | POST /api/goals/:id/transfer | Transfer money to or from a savings goal |
| GoalApi | Withdraw | POST /api/goals/:id/withdraw | Withdraw money from a savings goal |

## License

This project is licensed under the MIT License. See the [LICENSE] file for more details.

