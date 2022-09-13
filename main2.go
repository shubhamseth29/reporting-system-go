package main

// import (
// 	"context"
// 	"io/ioutil"
// 	"log"
// 	"sync"

// 	"github.com/chromedp/cdproto/page"
// 	"github.com/chromedp/chromedp"
// )

// func main() {
// 	ctx, cancel := chromedp.NewContext(context.Background(),
// 		// log CDP messages to understand the events better
// 		chromedp.WithDebugf(log.Printf),
// 	)
// 	defer cancel()

// 	// construct your html
// 	html := `<html lang="en">
// 	<head>
// 		<meta charset="UTF-8" />
// 		<meta http-equiv="X-UA-Compatible" content="IE=edge" />
// 		<meta name="viewport" content="width=device-width, initial-scale=1.0" />
// 		<title>Pie Portlet</title>
// 		<style>
// 			html {
// 				-webkit-print-color-adjust: exact;
// 			}
// 			html,
// 			body {
// 				margin: 0;
// 				padding: 0;
// 				box-sizing: border-box;
// 				color: #444444;
// 				background: white;
// 				font-family: helvetica;
// 				overflow-x: hidden;
// 			}
// 			thead {
// 				display: table-header-group;
// 				break-inside: avoid;
// 				page-break-inside: avoid;
// 			}
// 			table {
// 				width: 100%;
// 				border-collapse: collapse;
// 			}
// 			thead tr {
// 				border: 0;
// 				background: #263550;
// 				color: white;
// 			}
// 			tr:last-child {
// 			}
// 			tr:nth-child(even) {
// 				background: #f4f5f7;
// 			}
// 			tr {
// 				display: flex;
// 				border: 0.5px solid #c4c4c4;
// 				border-top: 0;
// 				break-inside: avoid;
// 				page-break-inside: avoid;
// 			}
// 			th {
// 				font-weight: 600;
// 				font-size: 12px;
// 				word-break: break-word;
// 			}

// 			th,
// 			td {
// 				flex: 1;
// 				padding: 8px;
// 				text-align: left;
// 				border-right: 1px solid #c4c4c4;
// 			}

// 			th:last-child,
// 			td:last-child {
// 				border-right: 0;
// 			}

// 			td {
// 				font-weight: 400;
// 				font-size: 10px;
// 				word-break: break-all;
// 			}
// 			.extra-message {
// 				width: 100%;
// 				text-align: center;
// 				margin-top: 50px;
// 				font-size: 1.5rem;
// 				color: #9a9999;
// 				display: none;
// 			}

// 			.description-container {
// 				margin: 20px 50px 0;
// 				display: none;
// 			}
// 			.description-label {
// 				display: block;
// 				margin-bottom: 8px;
// 				font-size: 1.2rem;
// 				font-weight: bold;
// 			}
// 			.description-content {
// 				font-size: 1rem;
// 				line-height: 1.4rem;
// 				text-align: justify;
// 				text-justify: inter-word;
// 			}
// 			.pdf-footer {
// 				height: 20px;
// 				position: absolute;
// 				bottom: 0;
// 				width: 100%;
// 				background: #4465a2;
// 				color: white;
// 				font-size: 0.7rem;
// 				display: flex;
// 				align-items: center;
// 				padding-left: 30px;
// 			}
// 		</style>
// 		<style>
// 			.pie-header {
// 				display: flex;
// 				flex-direction: column;
// 				width: 100%;
// 				margin: 0;
// 				padding: 0 50px 0;
// 				box-sizing: border-box;
// 			}
// 			.pie-title {
// 				font-weight: 500;
// 				font-size: 30px;
// 				margin: 0;
// 				padding-left: 5px;
// 			}
// 			.pie-header-bottom {
// 				width: 100%;
// 				height: 5px;
// 				margin-left: 50px;
// 				background: #263550;
// 			}
// 			.pie-container {
// 				display: flex;
// 				width: 11in;
// 				justify-content: center;
// 				align-items: center;
// 				page-break-after: auto;
// 			}
// 			.pie-content {
// 				display: flex;
// 				width: 9.9in;
// 				flex-direction: column;
// 				justify-content: center;
// 				align-items: center;
// 			}
// 			.pie-top-legends {
// 				width: 100%;
// 			}
// 			.pie-chart-div {
// 				width: 80%;
// 				height: 350px;
// 				padding-top: 20px;
// 				padding-bottom: 20px;
// 			}
// 			.second-pie-chart-div {
// 				width: 10%;
// 				height: 20px;
// 				display: none;
// 			}
// 			.pie-general {
// 				margin-bottom: 15px;
// 			}
// 			.total-cost-usage-key {
// 				font-weight: bold;
// 				font-size: 20px;
// 			}
// 			.total-cost-usage-value {
// 				font-size: 20px;
// 			}
// 			#pie-legendDivLine {
// 				width: 80%;
// 				display: none;
// 			}
// 			#pie-second-legendDivLine {
// 				width: 80%;
// 				height: 200px;
// 			}
// 			.pie-legendwrapper {
// 				width: 100%;
// 				display: flex;
// 				justify-content: center;
// 				align-items: center;
// 				margin-top: 20px;
// 			}
// 			.pie-table-container {
// 				margin: 20px 50px 0 50px;
// 			}
// 			.legendTd {
// 			}
// 			.legendTd div {
// 				height: 14px;
// 				width: 14px;
// 				border-radius: 100%;
// 				background-color: grey;
// 			}
// 		</style>
// 		<script
// 			defer
// 			src="https://code.jquery.com/jquery-3.2.1.min.js"
// 		></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/version/4.5.3/core.js"
// 		></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/version/4.5.3/charts.js"
// 		></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/4/themes/material.js"
// 		></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/4/themes/dark.js"
// 		></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/4/themes/animated.js"
// 		></script>
// 		<script defer src="https://www.amcharts.com/lib/3/ammap.js"></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/3/maps/js/worldLow.js"
// 		></script>
// 		<script
// 			defer
// 			src="https://www.amcharts.com/lib/3/themes/light.js"
// 		></script>
// 	</head>
// 	<body>
// 		<div class="pie-header">
// 			<h1 class="pie-title"></h1>
// 		</div>
// 		<div class="pie-header-bottom"></div>
// 		<div class="description-container">
// 			<span class="description-label">Description</span>
// 			<span class="description-content"> </span>
// 		</div>
// 		<div class="pie-container">
// 			<div class="pie-content">
// 				<div class="pie-chart-div"></div>
// 				<div class="pie-general">
// 					<span class="total-cost-usage-key">Total Cost:</span>
// 					<span class="total-cost-usage-value"></span>
// 				</div>
// 				<div id="pie-legendDivLine"></div>
// 				<div class="pie-legendwrapper">
// 					<div id="pie-second-legendDivLine"></div>
// 				</div>
// 			</div>
// 		</div>
// 		<div class="pie-table-container"></div>
// 		<div class="extra-message">No Data Available</div>
// 		<script defer>
// 			document.addEventListener('readystatechange', (event) => {
// 				if (event.target.readyState === 'complete') {
// 					(async () => {
// 						let chart;
// 						let pieSeries;
// 						let legends;
// 						let markerTemplate;
// 						let marker;
// 						let legendContainer;
// 						let newLegendContainer;
// 						let newLegend;
// 						let data = {
// 							next: {},
// 							previous: {},
// 							generalException: [],
// 							module: 'costmonitoring',
// 							insufficientPermission: [],
// 							invalidCredentials: [],
// 							referS3UrlBydefault: false,
// 							restrictedFilters: [],
// 							cloud: 'aws',
// 							resourceTags: [],
// 							insight: 'cost-by-service',
// 							noCredentialsDB: [],
// 							dataMap: {
// 								currencySymbol: '$',
// 								roundOff: 2,
// 								pie: { x: ['Cost($)'], y: ['Service'] },
// 								table: [
// 									{
// 										'Cost($)': 108948.17,
// 										Service: 'Amazon Elastic Compute Cloud'
// 									},
// 									{
// 										'Cost($)': 24502.15,
// 										Service:
// 											'Amazon Relational Database Service'
// 									},
// 									{
// 										'Cost($)': 8533.69,
// 										Service: 'Amazon Simple Storage Service'
// 									},
// 									{
// 										'Cost($)': 7429.98,
// 										Service: 'Elastic Load Balancing'
// 									},
// 									{
// 										'Cost($)': 6206.99,
// 										Service: 'Amazon CloudFront'
// 									},
// 									{
// 										'Cost($)': 5647.25,
// 										Service: 'Amazon DynamoDB'
// 									},
// 									{
// 										'Cost($)': 4932.45,
// 										Service: 'AmazonCloudWatch'
// 									},
// 									{
// 										'Cost($)': 4371.15,
// 										Service: 'Amazon ElastiCache'
// 									},
// 									{
// 										'Cost($)': 4088.43,
// 										Service: 'Amazon Redshift'
// 									},
// 									{
// 										'Cost($)': 3526.4,
// 										Service: 'Amazon API Gateway'
// 									},
// 									{
// 										'Cost($)': 3470.24,
// 										Service: 'Amazon Elastic File System'
// 									},
// 									{
// 										'Cost($)': 3364.95,
// 										Service: 'Savings Plans'
// 									},
// 									{
// 										'Cost($)': 2886.8,
// 										Service: 'Amazon Virtual Private Cloud'
// 									},
// 									{
// 										'Cost($)': 2673.81,
// 										Service: 'AWS CloudTrail'
// 									},
// 									{
// 										'Cost($)': 2390.05,
// 										Service: 'AWS Shield'
// 									},
// 									{
// 										'Cost($)': 1477.84,
// 										Service: 'AWS Key Management Service'
// 									},
// 									{
// 										'Cost($)': 1437.83,
// 										Service: 'AWS Network Firewall'
// 									},
// 									{
// 										'Cost($)': 1426.28,
// 										Service:
// 											'SUSE Linux Enterprise Server for SAP Applications 12 SP3'
// 									},
// 									{
// 										'Cost($)': 1278.2,
// 										Service: 'Amazon SageMaker'
// 									},
// 									{
// 										'Cost($)': 1214.45,
// 										Service:
// 											'Amazon Managed Streaming for Apache Kafka'
// 									},
// 									{
// 										'Cost($)': 1103.64,
// 										Service:
// 											'AWS Database Migration Service'
// 									},
// 									{
// 										'Cost($)': 981.33,
// 										Service:
// 											'Amazon Elastic Container Service'
// 									},
// 									{ 'Cost($)': 960.44, Service: 'AWS WAF' },
// 									{
// 										'Cost($)': 700.05,
// 										Service:
// 											'Amazon Elastic Container Service for Kubernetes'
// 									},
// 									{
// 										'Cost($)': 695.55,
// 										Service: 'AWS Lambda'
// 									},
// 									{
// 										'Cost($)': 633.78,
// 										Service: 'Amazon OpenSearch Service'
// 									},
// 									{
// 										'Cost($)': 590.43,
// 										Service: 'Amazon GuardDuty'
// 									},
// 									{
// 										'Cost($)': 567.81,
// 										Service: 'Amazon Kinesis Firehose'
// 									},
// 									{
// 										'Cost($)': 518.09,
// 										Service: 'AWS Device Farm'
// 									},
// 									{
// 										'Cost($)': 471.78,
// 										Service: 'Amazon Simple Email Service'
// 									},
// 									{ 'Cost($)': 439.24, Service: 'Amazon MQ' },
// 									{ 'Cost($)': 426.22, Service: 'AWS Glue' },
// 									{
// 										'Cost($)': 379.09,
// 										Service: 'Amazon QuickSight'
// 									},
// 									{
// 										'Cost($)': 357.87,
// 										Service: 'AWS Config'
// 									},
// 									{
// 										'Cost($)': 351.46,
// 										Service: 'AWS Directory Service'
// 									},
// 									{
// 										'Cost($)': 317.19,
// 										Service:
// 											'Amazon DocumentDB (with MongoDB compatibility)'
// 									},
// 									{
// 										'Cost($)': 274.49,
// 										Service: 'Amazon WorkSpaces'
// 									},
// 									{
// 										'Cost($)': 274.2,
// 										Service: 'AWS Transfer Family'
// 									},
// 									{
// 										'Cost($)': 241.5,
// 										Service:
// 											'Savings Plans for AWS Machine Learning'
// 									},
// 									{
// 										'Cost($)': 234.52,
// 										Service:
// 											'Amazon Managed Workflows for Apache Airflow'
// 									},
// 									{ 'Cost($)': 226.25, Service: 'AWS X-Ray' },
// 									{
// 										'Cost($)': 214.89,
// 										Service: 'Amazon Neptune'
// 									},
// 									{
// 										'Cost($)': 213.9,
// 										Service:
// 											'Palo Alto Networks VM-300 Bundle 2'
// 									},
// 									{
// 										'Cost($)': 182.87,
// 										Service: 'Amazon Polly'
// 									},
// 									{
// 										'Cost($)': 159.3,
// 										Service: 'Amazon CloudSearch'
// 									},
// 									{
// 										'Cost($)': 147.21,
// 										Service: 'Amazon Lightsail'
// 									},
// 									{
// 										'Cost($)': 132,
// 										Service: 'AWS Storage Gateway'
// 									},
// 									{
// 										'Cost($)': 128.86,
// 										Service: 'Amazon Simple Queue Service'
// 									},
// 									{
// 										'Cost($)': 123.09,
// 										Service: 'Amazon FSx'
// 									},
// 									{
// 										'Cost($)': 117.81,
// 										Service: 'DynamoDB Accelerator (DAX)'
// 									},
// 									{
// 										'Cost($)': 116.11,
// 										Service: 'AWS Support (Business)'
// 									},
// 									{
// 										'Cost($)': 112.95,
// 										Service:
// 											'Amazon EC2 Container Registry (ECR)'
// 									},
// 									{
// 										'Cost($)': 102.84,
// 										Service:
// 											'Oracle Enterprise Linux 7.2  supported by Navisite'
// 									},
// 									{
// 										'Cost($)': 101.54,
// 										Service: 'AWS Global Accelerator'
// 									},
// 									{
// 										'Cost($)': 91.2,
// 										Service:
// 											'Red Hat Enterprise Linux (RHEL) 7.9 with support by ProComputers'
// 									},
// 									{
// 										'Cost($)': 79.09,
// 										Service: 'Amazon AppStream'
// 									},
// 									{
// 										'Cost($)': 77,
// 										Service:
// 											'OpenVPN Access Server (10 Connected Devices)'
// 									},
// 									{
// 										'Cost($)': 73.96,
// 										Service: 'AmazonWorkMail'
// 									},
// 									{
// 										'Cost($)': 71.84,
// 										Service: 'Amazon Kinesis'
// 									},
// 									{
// 										'Cost($)': 67.54,
// 										Service: 'AWS Security Hub'
// 									},
// 									{
// 										'Cost($)': 61.85,
// 										Service:
// 											'Heimdall Proxy Standard Edition'
// 									},
// 									{
// 										'Cost($)': 61.8,
// 										Service:
// 											'OpenVPN Access Server (25 Connected Devices)'
// 									},
// 									{
// 										'Cost($)': 61.14,
// 										Service: 'AWS Secrets Manager'
// 									},
// 									{
// 										'Cost($)': 60.11,
// 										Service: 'Amazon Route 53'
// 									},
// 									{
// 										'Cost($)': 53.4,
// 										Service:
// 											'CloudEndure Disaster Recovery to AWS'
// 									},
// 									{
// 										'Cost($)': 53.29,
// 										Service:
// 											'Fortinet FortiGate Next-Generation Firewall'
// 									},
// 									{
// 										'Cost($)': 50.05,
// 										Service: 'Trend Micro Cloud One'
// 									},
// 									{
// 										'Cost($)': 34.88,
// 										Service:
// 											'SUSE Linux Enterprise Server for SAP Applications 15 SP3'
// 									},
// 									{
// 										'Cost($)': 28.32,
// 										Service: 'AWS Direct Connect'
// 									},
// 									{
// 										'Cost($)': 26,
// 										Service: 'AWS CodePipeline'
// 									},
// 									{
// 										'Cost($)': 26,
// 										Service: 'Amazon Registrar'
// 									},
// 									{ 'Cost($)': 24.45, Service: 'AWS Backup' },
// 									{
// 										'Cost($)': 23.88,
// 										Service:
// 											'Amazon Simple Notification Service'
// 									},
// 									{
// 										'Cost($)': 23.59,
// 										Service: 'Amazon WorkDocs'
// 									},
// 									{
// 										'Cost($)': 22.63,
// 										Service:
// 											'Plesk WordPress Edition PREMIUM, Website & WordPress Hosting'
// 									},
// 									{
// 										'Cost($)': 21.88,
// 										Service: 'AWS Amplify'
// 									},
// 									{
// 										'Cost($)': 19.36,
// 										Service:
// 											'CIS Amazon Linux 2 Benchmark - Level 2'
// 									},
// 									{
// 										'Cost($)': 17.96,
// 										Service: 'Amazon Inspector'
// 									},
// 									{
// 										'Cost($)': 15.5,
// 										Service:
// 											'Debian 10 (Debian Buster) with Support by Supported Images'
// 									},
// 									{
// 										'Cost($)': 15.4,
// 										Service:
// 											'Ubuntu 20 (Ubuntu 20.04 LTS) with Support by Supported Images'
// 									},
// 									{
// 										'Cost($)': 14.56,
// 										Service:
// 											'CloudEndure Disaster Recovery to AWS'
// 									},
// 									{
// 										'Cost($)': 13.3,
// 										Service:
// 											'CIS Microsoft Windows Server 2019 Benchmark - Level 2'
// 									},
// 									{
// 										'Cost($)': 12.46,
// 										Service:
// 											'Plesk Obsidian on Ubuntu - WordPress & Website Hosting Environment'
// 									},
// 									{
// 										'Cost($)': 11.48,
// 										Service:
// 											'WordPress Protected with End to End AppCalcium Security Solution'
// 									},
// 									{
// 										'Cost($)': 11.47,
// 										Service:
// 											'Fortinet Managed Rules for AWS WAF Classic - Complete OWASP Top 10'
// 									},
// 									{
// 										'Cost($)': 10.85,
// 										Service:
// 											'Plesk Onyx on Windows 2012 R2 WordPress & Website Hosting Environment'
// 									},
// 									{
// 										'Cost($)': 10.78,
// 										Service:
// 											'OpenVPN Access Server (5 Connected Devices)'
// 									},
// 									{
// 										'Cost($)': 9.3,
// 										Service: 'Amazon Inspector'
// 									},
// 									{
// 										'Cost($)': 9.12,
// 										Service: 'Amazon Fraud Detector'
// 									},
// 									{ 'Cost($)': 6, Service: 'Amazon Athena' },
// 									{
// 										'Cost($)': 4.17,
// 										Service: 'Amazon Textract'
// 									},
// 									{
// 										'Cost($)': 3.55,
// 										Service: 'AWS Cost Explorer'
// 									},
// 									{ 'Cost($)': 3.29, Service: 'CodeBuild' },
// 									{
// 										'Cost($)': 2.45,
// 										Service: 'Amazon Glacier'
// 									},
// 									{
// 										'Cost($)': 1.66,
// 										Service: 'Amazon Macie'
// 									},
// 									{
// 										'Cost($)': 1.63,
// 										Service:
// 											'Debian GNU/Linux 9 (Stretch) with support by Frontline'
// 									},
// 									{
// 										'Cost($)': 1.62,
// 										Service: 'AWS Step Functions'
// 									},
// 									{
// 										'Cost($)': 1.49,
// 										Service: 'Amazon Pinpoint'
// 									},
// 									{ 'Cost($)': 1.19, Service: 'AWS IoT' },
// 									{
// 										'Cost($)': 0.57,
// 										Service: 'Amazon Rekognition'
// 									},
// 									{
// 										'Cost($)': 0.4,
// 										Service:
// 											'CentOS 7 with support by ProComputers'
// 									},
// 									{ 'Cost($)': 0.37, Service: 'AWS AppSync' },
// 									{
// 										'Cost($)': 0.28,
// 										Service: 'AWS Systems Manager'
// 									},
// 									{
// 										'Cost($)': 0.18,
// 										Service:
// 											'Amazon Keyspaces (for Apache Cassandra)'
// 									},
// 									{
// 										'Cost($)': 0.16,
// 										Service: 'Amazon Cognito'
// 									},
// 									{
// 										'Cost($)': 0.15,
// 										Service: 'InfluxDB Cloud'
// 									},
// 									{
// 										'Cost($)': 0.13,
// 										Service: 'Amazon AppFlow'
// 									},
// 									{
// 										'Cost($)': 0.1,
// 										Service:
// 											'Red Hat Enterprise Linux (RHEL) 7.5 with support by ProComputers'
// 									},
// 									{
// 										'Cost($)': 0.09,
// 										Service: 'AWS Compute Optimizer'
// 									},
// 									{
// 										'Cost($)': 0.08,
// 										Service: 'Amazon Detective'
// 									},
// 									{
// 										'Cost($)': 0.07,
// 										Service: 'CloudWatch Events'
// 									},
// 									{ 'Cost($)': 0.01, Service: 'AWS Budgets' },
// 									{
// 										'Cost($)': 0,
// 										Service: 'Amazon Comprehend'
// 									},
// 									{ 'Cost($)': 0, Service: 'Amazon Lex' },
// 									{
// 										'Cost($)': 0,
// 										Service: 'Amazon SimpleDB'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'AWS Data Exchange'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'Amazon Kinesis Analytics'
// 									},
// 									{ 'Cost($)': 0, Service: 'AWS DataSync' },
// 									{ 'Cost($)': 0, Service: 'AWS CloudShell' },
// 									{
// 										'Cost($)': 0,
// 										Service: 'AWS CodeArtifact'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'Amazon Timestream'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'AWS Service Catalog'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service:
// 											'Contact Center Telecommunications (service sold by AMCS, LLC) '
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'Amazon Sumerian'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service:
// 											'Amazon Elastic Container Registry Public'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service:
// 											'VM-Series Next-Generation Firewall Bundle 2'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'Amazon Elastic Transcoder'
// 									},
// 									{ 'Cost($)': 0, Service: 'AWSCloudShell' },
// 									{
// 										'Cost($)': 0,
// 										Service: 'AmazonTimestream'
// 									},
// 									{
// 										'Cost($)': 0,
// 										Service: 'AWSCodeArtifact'
// 									},
// 									{ 'Cost($)': 0, Service: 'AmazonSimpleDB' },
// 									{ 'Cost($)': 0, Service: 'AmazonStates' },
// 									{ 'Cost($)': 0, Service: 'AWSDataSync' },
// 									{
// 										'Cost($)': 0,
// 										Service: 'AmazonKinesisAnalytics'
// 									},
// 									{ 'Cost($)': 0, Service: 'AmazonLex' },
// 									{
// 										'Cost($)': 0,
// 										Service: 'AWSSystemsManager'
// 									},
// 									{ 'Cost($)': 0, Service: 'AWSEvents' },
// 									{ 'Cost($)': 0, Service: 'comprehend' },
// 									{
// 										'Cost($)': -0.01,
// 										Service: 'AmazonDetective'
// 									},
// 									{
// 										'Cost($)': -0.01,
// 										Service: 'AmazonCognito'
// 									},
// 									{ 'Cost($)': -0.01, Service: 'AmazonMCS' },
// 									{ 'Cost($)': -0.03, Service: 'AWSAppSync' },
// 									{
// 										'Cost($)': -0.04,
// 										Service: 'AmazonPinpoint'
// 									},
// 									{
// 										'Cost($)': -0.04,
// 										Service: 'AmazonRekognition'
// 									},
// 									{ 'Cost($)': -0.09, Service: 'AWSIoT' },
// 									{
// 										'Cost($)': -0.12,
// 										Service: 'AmazonMacie'
// 									},
// 									{
// 										'Cost($)': -0.12,
// 										Service: 'AmazonWorkMail'
// 									},
// 									{
// 										'Cost($)': -0.14,
// 										Service: 'AmazonTextract'
// 									},
// 									{
// 										'Cost($)': -0.17,
// 										Service: 'AmazonGlacier'
// 									},
// 									{
// 										'Cost($)': -0.23,
// 										Service: 'AWSCostExplorer'
// 									},
// 									{ 'Cost($)': -0.25, Service: 'CodeBuild' },
// 									{
// 										'Cost($)': -0.34,
// 										Service: 'AmazonAthena'
// 									},
// 									{ 'Cost($)': -0.38, Service: 'AWSGlue' },
// 									{
// 										'Cost($)': -0.59,
// 										Service: 'AmazonWorkDocs'
// 									},
// 									{
// 										'Cost($)': -0.67,
// 										Service: 'AmazonCloudSearch'
// 									},
// 									{
// 										'Cost($)': -1.26,
// 										Service: 'AmazonInspectorV2'
// 									},
// 									{ 'Cost($)': -1.53, Service: 'AWSAmplify' },
// 									{
// 										'Cost($)': -1.63,
// 										Service: 'AmazonLightsail'
// 									},
// 									{ 'Cost($)': -1.72, Service: 'AWSBackup' },
// 									{ 'Cost($)': -1.8, Service: 'AmazonSNS' },
// 									{
// 										'Cost($)': -1.82,
// 										Service: 'AWSCodePipeline'
// 									},
// 									{ 'Cost($)': -1.9, Service: 'AmazonFSx' },
// 									{
// 										'Cost($)': -2.32,
// 										Service: 'AWSSecurityHub'
// 									},
// 									{
// 										'Cost($)': -2.47,
// 										Service: 'AWSDirectConnect'
// 									},
// 									{
// 										'Cost($)': -2.59,
// 										Service: 'AmazonRoute53'
// 									},
// 									{
// 										'Cost($)': -2.86,
// 										Service: 'AmazonAppStream'
// 									},
// 									{
// 										'Cost($)': -4.17,
// 										Service: 'AWSSecretsManager'
// 									},
// 									{
// 										'Cost($)': -5.67,
// 										Service: 'AmazonKinesis'
// 									},
// 									{
// 										'Cost($)': -7.44,
// 										Service: 'AWSStorageGateway'
// 									},
// 									{
// 										'Cost($)': -8.44,
// 										Service: 'AWSQueueService'
// 									},
// 									{ 'Cost($)': -8.83, Service: 'AmazonECR' },
// 									{
// 										'Cost($)': -10.01,
// 										Service: 'AWSGlobalAccelerator'
// 									},
// 									{
// 										'Cost($)': -12.67,
// 										Service: 'AWSDirectoryService'
// 									},
// 									{
// 										'Cost($)': -12.8,
// 										Service: 'AmazonPolly'
// 									},
// 									{ 'Cost($)': -15.84, Service: 'AWSXRay' },
// 									{
// 										'Cost($)': -21.5,
// 										Service: 'AWSTransfer'
// 									},
// 									{ 'Cost($)': -21.66, Service: 'AmazonES' },
// 									{
// 										'Cost($)': -25.79,
// 										Service: 'AmazonNeptune'
// 									},
// 									{ 'Cost($)': -33.22, Service: 'AWSConfig' },
// 									{ 'Cost($)': -36.31, Service: 'AmazonEKS' },
// 									{
// 										'Cost($)': -38.06,
// 										Service: 'AmazonDocDB'
// 									},
// 									{
// 										'Cost($)': -39.73,
// 										Service: 'AmazonKinesisFirehose'
// 									},
// 									{
// 										'Cost($)': -40.9,
// 										Service: 'AmazonWorkSpaces'
// 									},
// 									{
// 										'Cost($)': -41.34,
// 										Service: 'AmazonQuickSight'
// 									},
// 									{
// 										'Cost($)': -43.15,
// 										Service: 'AmazonGuardDuty'
// 									},
// 									{ 'Cost($)': -48.66, Service: 'AmazonMQ' },
// 									{ 'Cost($)': -50.14, Service: 'AWSLambda' },
// 									{ 'Cost($)': -50.87, Service: 'AmazonSES' },
// 									{
// 										'Cost($)': -62.17,
// 										Service: 'AWSDeviceFarm'
// 									},
// 									{ 'Cost($)': -63.16, Service: 'awswaf' },
// 									{ 'Cost($)': -64.38, Service: 'AmazonECS' },
// 									{
// 										'Cost($)': -91.59,
// 										Service: 'AmazonSageMaker'
// 									},
// 									{ 'Cost($)': -92.56, Service: 'AmazonMSK' },
// 									{ 'Cost($)': -101.11, Service: 'awskms' },
// 									{
// 										'Cost($)': -133.27,
// 										Service: 'AmazonVPC'
// 									},
// 									{
// 										'Cost($)': -161.8,
// 										Service: 'AWSDatabaseMigrationSvc'
// 									},
// 									{ 'Cost($)': -167.3, Service: 'AWSShield' },
// 									{
// 										'Cost($)': -187.9,
// 										Service: 'AWSCloudTrail'
// 									},
// 									{
// 										'Cost($)': -207.18,
// 										Service: 'ComputeSavingsPlans'
// 									},
// 									{
// 										'Cost($)': -209.01,
// 										Service: 'AWSNetworkFirewall'
// 									},
// 									{
// 										'Cost($)': -247.4,
// 										Service: 'AmazonApiGateway'
// 									},
// 									{
// 										'Cost($)': -292.22,
// 										Service: 'AmazonRedshift'
// 									},
// 									{
// 										'Cost($)': -305.3,
// 										Service: 'AmazonCloudWatch'
// 									},
// 									{
// 										'Cost($)': -320.83,
// 										Service: 'AmazonElastiCache'
// 									},
// 									{
// 										'Cost($)': -326.52,
// 										Service: 'AmazonDynamoDB'
// 									},
// 									{ 'Cost($)': -472.78, Service: 'AmazonS3' },
// 									{ 'Cost($)': -542.7, Service: 'AmazonEFS' },
// 									{ 'Cost($)': -572.68, Service: 'AWSELB' },
// 									{
// 										'Cost($)': -894.5,
// 										Service: 'AmazonCloudFront'
// 									},
// 									{
// 										'Cost($)': -1824.79,
// 										Service: 'AmazonRDS'
// 									},
// 									{
// 										'Cost($)': -8063.56,
// 										Service: 'AmazonEC2'
// 									}
// 								],
// 								totalCost: 197693.44,
// 								tableKeys: ['Service', 'Cost($)']
// 							},
// 							dataList: [],
// 							otherException: [],
// 							page: 'home',
// 							insightText: 'Cost of AWS Services',
// 							startDate: '14 June 2022',
// 							endDate: '21 June 2022'
// 						};
// 						let showGraph = true;
// 						let legendsData = [];
// 						let legendsColorObj = {};
// 						$('.pie-title').attr(
// 							'id',
// 							data['cloud'] +
// 								data['module'] +
// 								data['page'] +
// 								data['insight'] +
// 								'insightText'
// 						);
// 						$('.pie-chart-div').attr(
// 							'id',
// 							data['cloud'] +
// 								data['module'] +
// 								data['page'] +
// 								data['insight']
// 						);
// 						$('.pie-table-container').attr(
// 							'id',
// 							data['cloud'] +
// 								data['module'] +
// 								data['page'] +
// 								data['insight'] +
// 								'table'
// 						);
// 						//$('.pie-legendDivLine').attr('id', data['cloud'] + data['module'] + data['page'] + data['insight'] + 'legenddiv');
// 						$(
// 							'#' +
// 								data['cloud'] +
// 								data['module'] +
// 								data['page'] +
// 								data['insight'] +
// 								'insightText'
// 						).html(data['insightText']);

// 						if (data['dataMap']['totalCost']) {
// 							$('.total-cost-usage-key').html('Total Cost:');
// 							$('.total-cost-usage-value').html(
// 								data['dataMap']['currencySymbol'] +
// 									data['dataMap']['totalCost']
// 							);
// 						} else if (data['dataMap']['totalUsage']) {
// 							$('.total-cost-usage-key').html('Total Usage:');
// 							$('.total-cost-usage-value').html(
// 								data['dataMap']['totalUsage']
// 							);
// 						}

// 						if (
// 							data['checkDescription'] &&
// 							data['checkDescription']['description']
// 						) {
// 							$('.description-container').css({
// 								display: 'block'
// 							});
// 							$('.description-content').html(
// 								data['checkDescription']['description']
// 							);
// 							$('.pie-container').css({ 'margin-top': '0' });
// 							$('.extra-message').css({ 'margin-top': '20px' });
// 						}

// 						if (Object.keys(data['dataMap']).length === 0) {
// 							$('.extra-message').css({ display: 'block' });
// 							$('.pie-container').css({ display: 'none' });
// 							$('.pie-table-container').css({ display: 'none' });
// 							return;
// 						}

// 						// prepare pie data
// 						const tableDataAfterColSel = [];
// 						const tableDataAfterColSelNew = [];
// 						const prepareTempData = {};
// 						data['dataMap']['table'].forEach((element) => {
// 							if (
// 								element[data['dataMap']['pie']['y'][0]] in
// 								prepareTempData
// 							) {
// 								prepareTempData[
// 									element[data['dataMap']['pie']['y'][0]]
// 								] = +(
// 									prepareTempData[
// 										element[data['dataMap']['pie']['y'][0]]
// 									] + element[data['dataMap']['pie']['x'][0]]
// 								).toFixed(data['dataMap']['roundOff']);
// 							} else {
// 								prepareTempData[
// 									element[data['dataMap']['pie']['y'][0]]
// 								] = element[data['dataMap']['pie']['x'][0]];
// 							}
// 						});

// 						Object.keys(prepareTempData).forEach((key) => {
// 							if (!(prepareTempData[key] < 0)) {
// 								const tempObject = {};
// 								tempObject['y'] = key;
// 								tempObject['x'] = prepareTempData[key];
// 								tableDataAfterColSel.push(tempObject);
// 								tableDataAfterColSelNew.push(tempObject);
// 							}
// 						});
// 						// prepare pie data

// 						if (
// 							('totalCost' in data['dataMap'] &&
// 								data['dataMap']['totalCost'] === 0) ||
// 							('totalUsage' in data['dataMap'] &&
// 								data['dataMap']['totalUsage'] === 0) ||
// 							tableDataAfterColSel.length > 200
// 						) {
// 							$('.pie-container').css({ display: 'none' });
// 							showGraph = false;
// 						}

// 						if (showGraph) {
// 							chart = am4core.create(
// 								data['cloud'] +
// 									data['module'] +
// 									data['page'] +
// 									data['insight'],
// 								am4charts.PieChart
// 							);
// 							chart.data = tableDataAfterColSel;

// 							pieSeries = chart.series.push(
// 								new am4charts.PieSeries()
// 							);
// 							pieSeries.dataFields.value = 'x';
// 							pieSeries.dataFields.category = 'y';
// 							pieSeries.slices.template.stroke =
// 								am4core.color('#e3e3e3');
// 							pieSeries.slices.template.strokeWidth = 0.5;
// 							pieSeries.slices.template.strokeOpacity = 1;
// 							pieSeries.slices.template.margin = 80;
// 							pieSeries.labels.template.disabled = true;
// 							pieSeries.ticks.template.disabled = true;
// 							// pieSeries.labels.template.adapter.add('text', (label, target, key) => {

// 							//     if (target.dataItem && (target.dataItem.value > tableDataAfterColSel[5][data['dataMap']['pie']['x'][0]])) {
// 							//       return label;
// 							//     }
// 							//     return '';
// 							//   });
// 							//pieSeries.legendSettings.itemValueText = '{valueY}';
// 							chart.radius = am4core.percent(80);

// 							pieSeries.hiddenState.properties.opacity = 0.5;
// 							pieSeries.hiddenState.properties.endAngle = -90;
// 							pieSeries.hiddenState.properties.startAngle = -90;

// 							legends = new am4charts.Legend();
// 							chart.legend = legends;
// 							legendContainer = am4core.create(
// 								'pie-legendDivLine',
// 								am4core.Container
// 							);
// 							chart.legend.parent = legendContainer;
// 							//legendContainer.width = am4core.percent(100);
// 							//legendContainer.height = am4core.percent(100);
// 							chart.legend.labels.template.maxWidth = 300;
// 							chart.legend.labels.template.truncate = false;
// 							chart.legend.labels.template.wrap = true;
// 							chart.legend.useDefaultMarker = true;
// 							markerTemplate = chart.legend.markers.template;
// 							markerTemplate.width = 15;
// 							markerTemplate.height = 15;
// 							chart.legend.fontSize = '11';
// 							marker =
// 								chart.legend.markers.template.children.getIndex(
// 									0
// 								);
// 							marker.cornerRadius(12, 12, 12, 12);
// 							marker.radius = 5.5;
// 							marker.stroke = am4core.color('#ccc');
// 						}

// 						setTimeout(() => {
// 							let chartRows;
// 							if (showGraph) {
// 								chartRows = legends['labels']['_values'];

// 								// code for hidden pie chart
// 								let tempData;

// 								if (tableDataAfterColSelNew.length > 4) {
// 									tableDataAfterColSelSorted =
// 										tableDataAfterColSelNew.sort((a, b) =>
// 											a.x < b.x ? 1 : -1
// 										);
// 									tempData =
// 										tableDataAfterColSelSorted.splice(
// 											0,
// 											10
// 										);
// 								} else {
// 									tempData = tableDataAfterColSelNew;
// 								}

// 								tempData.forEach((row) => {
// 									legendsData.push({
// 										name:
// 											row['y'] +
// 											' : ' +
// 											(data['dataMap']['currencySymbol']
// 												? data['dataMap'][
// 														'currencySymbol'
// 												  ]
// 												: '') +
// 											row['x'],
// 										value: row['x'],
// 										fill: getLegendColor(
// 											row,
// 											chartRows,
// 											'y'
// 										)
// 									});
// 								});
// 								newLegendContainer = am4core.create(
// 									'pie-second-legendDivLine',
// 									am4core.Container
// 								);
// 								newLegendContainer.width = am4core.percent(100);
// 								newLegend = new am4charts.Legend();
// 								newLegend.parent = newLegendContainer;
// 								newLegend.itemContainers.template.togglable = false;
// 								newLegend.marginTop = 20;
// 								newLegend.data = legendsData;
// 								newLegend.labels.template.maxWidth = 300;
// 								newLegend.labels.template.truncate = false;
// 								newLegend.labels.template.wrap = true;
// 								newLegend.useDefaultMarker = true;
// 								let newMarkerTemplate =
// 									newLegend.markers.template;
// 								newMarkerTemplate.width = 15;
// 								newMarkerTemplate.height = 15;
// 								newLegend.fontSize = '12';
// 								let newMarker =
// 									newLegend.markers.template.children.getIndex(
// 										0
// 									);
// 								newMarker.cornerRadius(12, 12, 12, 12);
// 								newMarker.radius = 6;
// 								newMarker.stroke = am4core.color('#ccc');

// 								setTimeout(() => {
// 									z;
// 									document.getElementById(
// 										'pie-second-legendDivLine'
// 									).style.height =
// 										newLegend.contentHeight + 'px';
// 								}, 10);

// 								console.log('legendsdata', newLegend.data);
// 							}

// 							let htmlForTable = '<table>';
// 							htmlForTable += '<thead><tr>';
// 							let keys = [];
// 							if (
// 								data['dataMap']['tableKeys'] &&
// 								data['dataMap']['tableKeys'].length !== 0
// 							) {
// 								keys = data['dataMap']['tableKeys'];
// 							} else {
// 								keys = Object.keys(data['dataMap']['table'][0]);
// 							}
// 							// if (showGraph) {
// 							//     htmlForTable += '<th>' + 'Legends' + '</th>';
// 							// }
// 							keys.forEach((key) => {
// 								htmlForTable += '<th>' + key + '</th>';
// 							});
// 							htmlForTable += '</tr></thead>';
// 							htmlForTable += '<tbody>';
// 							for (
// 								let i = 0,
// 									length = data['dataMap']['table'].length;
// 								i < length;
// 								++i
// 							) {
// 								htmlForTable += '<tr>';
// 								// if (showGraph) {
// 								//     // Lengend color
// 								//     htmlForTable += '<td class="legendTd">
// 								//     <div class="key-${i}">
// 								//     </div>
// 								//     </td>';

// 								//     setTimeout(() => {
// 								//         $('.key-' + i).css('background-color', getLegendColor(data['dataMap']['table'][i], chartRows, data['dataMap']['pie']['y'][0]));

// 								//         if (i === 0) {
// 								//             document.getElementById('pie-second-legendDivLine').style.height = newLegend.contentHeight ;
// 								//         }
// 								//     }, 10);
// 								// }
// 								keys.forEach((key) => {
// 									htmlForTable +=
// 										'<td>' +
// 										data['dataMap']['table'][i][key] +
// 										'</td>';
// 								});
// 								htmlForTable += '</tr>';
// 							}
// 							htmlForTable += '</tbody>';
// 							htmlForTable += '</table>';
// 							$(
// 								'#' +
// 									data['cloud'] +
// 									data['module'] +
// 									data['page'] +
// 									data['insight'] +
// 									'table'
// 							).html(htmlForTable);
// 						}, 3000);
// 					})();

// 					let getLegendColor = (cellItem, rowsData, categoryAxis) => {
// 						let color = '#ffffff';
// 						rowsData.forEach((rowData) => {
// 							if (
// 								rowData['_dataItem']['_dataContext']['y'] ===
// 								cellItem[categoryAxis]
// 							) {
// 								color =
// 									rowData['_dataItem']['legendDataItem'][
// 										'colorOrig'
// 									]['hex'];
// 							}
// 						});
// 						return color;
// 					};
// 				}
// 			});
// 		</script>
// 	</body>
// </html>

// `

// 	var wg sync.WaitGroup
// 	wg.Add(1)
// 	if err := chromedp.Run(ctx,
// 		chromedp.Navigate("about:blank"),
// 		// setup the listener to listen for the page.EventLoadEventFired
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			lctx, cancel := context.WithCancel(ctx)
// 			chromedp.ListenTarget(lctx, func(ev interface{}) {
// 				if _, ok := ev.(*page.EventLoadEventFired); ok {
// 					wg.Done()
// 					// remove event listener
// 					cancel()
// 				}
// 			})
// 			return nil
// 		}),
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			frameTree, err := page.GetFrameTree().Do(ctx)
// 			if err != nil {
// 				return err
// 			}
// 			return page.SetDocumentContent(frameTree.Frame.ID, html).Do(ctx)
// 		}),
// 		// wait for the page.EventLoadEventFired
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			wg.Wait()
// 			return nil
// 		}),
// 		chromedp.ActionFunc(func(ctx context.Context) error {
// 			buf, _, err := page.PrintToPDF().WithPrintBackground(false).Do(ctx)
// 			if err != nil {
// 				return err
// 			}
// 			return ioutil.WriteFile("sample.pdf", buf, 0644)
// 		}),
// 	); err != nil {
// 		log.Fatal(err)
// 	}
// }