##### Version v1.7.4 (2021-09-06)

	Added fields for recurring payments: recurring-frequency and recurring-expiry.

##### Version v1.7.3 (2021-06-04)

	Added error codes for soft declines

##### Version v1.7.2 (2020-09-02)

	Add merchant-transaction-id to payment response parsing
	Affected:
	 - payment response
	 - result response
	 - callback parsing

##### Version v1.7.1 (2020-08-05)

	Add parameters describing cardholder device

##### Version v1.7.0 (2020-07-03)

	Improve authorization to use digest instead of API key.
	Verify non-failed responses for valid digest.
	Add possibility to validate callback data.
	Implement /report method.
	Implement response parsing.

##### Version 1.6.0 (2020-02-25)

	Add possibility to use custom return URL

##### Version 1.5.0 (2019-07-17)

	Add tokenization feature

##### Version 1.4.0 (2019-05-09)

	Add card verification

##### Version 1.3.3 (2019-03-21)

	Add custom 3D return URL

##### Version 1.3.2 (2019-03-04)

	Add merchant-referring-name fields to an order.

##### Version 1.3.1 (2019-01-23)

	Fix methods DMS Hold Charge and Cancel: methods require merchant transaction ID.

##### Version 1.2.0 (2018-08-21)

	Added missingh methods B2P and card 3-D Secure enrollment verification

##### Version 1.1.1 (2018-02-27)

	Minor fixes

##### Version 1.1.0 (2018-01-22)

	Add routes for init recurrent SMS and DMS hold

##### Version 1.0.0 (2017-11-10)

	First release
