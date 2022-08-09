# Feegrant

## Oveview

This module allows accounts to grant fee allowances and to use fees from their accounts. Grantees can execute any transaction without the need to maintain sufficient fees.

In order to make blockchain transactions, the signing account must possess a sufficient balance of the right denomination in order to pay fees. There are classes of transactions where needing to maintain a wallet with sufficient fees is a barrier to adoption.

For instance, when proper permissions are setup, someone may temporarily delegate the ability to vote on proposals to a "burner" account that is stored on a mobile phone with only minimal security.

Other use cases include workers tracking items in a supply chain or farmers submitting field data for analytics or compliance purposes.

For all of these use cases, UX would be significantly enhanced by obviating the need for these accounts to always maintain the appropriate fee balance. This is especially true if we wanted to achieve enterprise adoption for something like supply chain tracking.

While one solution would be to have a service that fills up these accounts automatically with the appropriate fees, a better UX would be provided by allowing these accounts to pull from a common fee pool account with proper spending limits. A single pool would reduce the churn of making lots of small "fill up" transactions and also more effectively leverages the resources of the organization setting up the pool.

## Concepts