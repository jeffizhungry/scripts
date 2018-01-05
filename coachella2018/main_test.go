package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseQueueLink(t *testing.T) {
	testcases := map[string]struct {
		html     string
		expected string
	}{
		"should be able to find queue id link": {
			// From Summer 2017 Advanced Sale for Weekend 1
			html: `<!DOCTYPE html>
			<!-- saved from url=(0135)http://go.festivalticketing.com/?c=festivalticketing&e=coachellawknd12018a&q=7065486b-36d2-4d20-a06b-4f6d53580286&cid=en-US&l=Coachella -->
			<html lang="en"><head><meta http-equiv="Content-Type" content="text/html; charset=UTF-8">
			    
			<meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
			<meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">

			<meta id="queue-it_log" data-userid="66b37dfa-cddd-4f8d-947c-9860ab0ca29e" data-proxyurl="https://logging.queue-it.net/logging/event">
			<title>Coachella 2018 Queue</title>

			    
			    <link rel="apple-touch-icon-precomposed" sizes="144x144" href="http://assets.queue-it.net/static/QueueFront/img/apple-touch-icon-144x144_d72f69f877cca3aa3f8ddb61337a4714.png">
			<link rel="apple-touch-icon-precomposed" sizes="152x152" href="http://assets.queue-it.net/static/QueueFront/img/apple-touch-icon-152x152_f94062b2cee189af07d20c53acdb2629.png">
			<link rel="apple-touch-icon" href="http://assets.queue-it.net/static/QueueFront/img/queue-it_appletouch_d72f69f877cca3aa3f8ddb61337a4714.png">
			<link rel="icon" type="image/png" href="http://assets.queue-it.net/static/QueueFront/img/favicon-32x32_f6ab48d8aa41b9414536f887c807b5d2.png" sizes="32x32">
			<link rel="icon" type="image/png" href="http://assets.queue-it.net/static/QueueFront/img/favicon-16x16_64f4c2b4ed8ab849d54dd4008cfb2a50.png" sizes="16x16">
			<link rel="shortcut icon" href="http://assets.queue-it.net/festivalticketing/userdata/coach/wk1/favicon.ico">
			<meta name="application-name" content=" ">
			<meta name="msapplication-TileColor" content="#FFFFFF">
			<meta name="msapplication-TileImage" content="//assets.queue-it.net/static/QueueFront/img/mstile-144x144_d72f69f877cca3aa3f8ddb61337a4714.png">
			    <link rel="stylesheet" href="./Coachella 2018 Queue_files/style_161681ed0ebcbf83a2b826053fc8aa5c.css" type="text/css"><!--[if lte IE 7]><link rel='stylesheet' href='//assets.queue-it.net/static/QueueFront/css/combined/ie_27516bd2e87d5dcd97fc4859fa766d2d.css' type='text/css' /><![endif]--><link rel="stylesheet" href="./Coachella 2018 Queue_files/coachW1w.css" type="text/css">

			    <iframe height="0" width="0" src="./Coachella 2018 Queue_files/activityi.html" style="display: none; visibility: hidden;"></iframe><script src="./Coachella 2018 Queue_files/217011611820041" async=""></script><script src="./Coachella 2018 Queue_files/291582177684507" async=""></script><script async="" src="./Coachella 2018 Queue_files/fbevents.js"></script><script type="text/javascript" async="" src="./Coachella 2018 Queue_files/conversion_async.js"></script><script type="text/javascript" async="" src="./Coachella 2018 Queue_files/conversion_async.js"></script><script type="text/javascript" async="" src="./Coachella 2018 Queue_files/conversion_async.js"></script><script async="" src="./Coachella 2018 Queue_files/cf.js"></script><script async="" src="./Coachella 2018 Queue_files/gtm.js"></script><script async="" src="./Coachella 2018 Queue_files/universalpixel.js"></script><script>
			    (function () {
			        if (!window.console)
			            window.console = {};

			        if (!window.console.log)
			            window.console.log = function () { };

			        if (!window.console.debug)
			            window.console.debug = function () { };
			    })();
			</script>
			<script>
			    function getScriptBasePath() {
			        return '/Script/combined/';
			    }
			</script>
			<script type="text/javascript" src="./Coachella 2018 Queue_files/common_9a864e9d04f3211883ddca8e38ceb8e6.js"></script>
			<script>window.jQuery || document.write('<script src="' + getScriptBasePath() + 'common.js' + '"><\/script>')</script>
			<script type="text/javascript" src="./Coachella 2018 Queue_files/queueit_716f646e3c1a61f893d383b49646d4a0.js"></script>
			<script>
			    window.QueueIt || document.write('<script src="' + getScriptBasePath() + 'queueit.js' + '"><\/script>');

			    if (window.soundManager)
			        window.soundManager.debugMode = false;
			</script>
			    <script type="text/javascript" id="queueit-perftiming-script" src="./Coachella 2018 Queue_files/statistics.min.js" data-queueit-c="festivalticketing" data-queueit-host="eu-west-1-perf-api.queue-it.net" data-queueit-tag-eventid="coachellawknd12018a" data-queueit-tag-queueid="7065486b-36d2-4d20-a06b-4f6d53580286" data-queueit-tag-queueit="queue">
			    </script>



			    <script type="text/javascript">

			    window.queueViewModel = new QueueIt.Queue.InQueueView({
			        customerId: 'festivalticketing',
			        eventId: 'coachellawknd12018a',
			        queueId: '7065486b-36d2-4d20-a06b-4f6d53580286',
			        targetUrl: decodeURIComponent(''),
			        customUrlParams: '',
			        culture: 'en-US',
			        layout: 'Coachella',
			        layoutVersion:'1496272606',
			        messageFeed: window.messageFeed,
			        emailAddress: '',
			        sdkVersion: '',
			        isBeforeOrIdle: true,
			        inqueueInfo: {"isBeforeOrIdle":true,"pageId":"before","pageClass":"before prequeue","QueueState":1,"ticket":{"highLoad":null,"queuePaused":null,"lastUpdatedUTC":"\/Date(1496422823561)\/","queueNumber":null,"whichIsIn":null,"expectedServiceTime":null,"usersInLineAheadOfYou":null,"lastUpdated":"10:00:23 AM","progress":null,"timeZonePostfix":"PST","expectedServiceTimeUTC":"\/Date(-62135596800000)\/","customUrlParams":"","sdkVersion":null,"windowStartTimeUTC":null,"windowStartTime":"","secondsToStart":3636,"usersInQueue":0,"eventStartTimeFormatted":"11:00 AM","eventStartTimeUTC":"\/Date(1496426400000)\/","QueueId":"7065486b-36d2-4d20-a06b-4f6d53580286","IsEventIdle":false},"message":null,"texts":{"countdownFinishedText":"You will now be redirected to the line.","header":"Coachella Weekend 1","body":"<p>The Coachella Weekend 1 advance sale has not yet begun. When the advance sale begins you will automatically be assigned a random place in line. You will have 10 minutes to complete your purchase after you enter the shopping cart.</p><p>For details on passes that are available please <a href=\"https://www.coachella.com/coachella-2018-advanced-sale-information/\" target=\"_blank\">click here</a>.</p><p>You can review the step by step purchase process <a href=\"https://www.coachella.com/guidebook/how-to-purchase/\" target=\"_blank\">here</a>.</p>","disclaimerText":"","styleSheets":"<link rel='stylesheet' href='//assets.queue-it.net/static/QueueFront/css/combined/style_161681ed0ebcbf83a2b826053fc8aa5c.css' type='text/css' /><!--[if lte IE 7]><link rel='stylesheet' href='//assets.queue-it.net/static/QueueFront/css/combined/ie_27516bd2e87d5dcd97fc4859fa766d2d.css' type='text/css' /><![endif]--><link rel='stylesheet' href='//assets.queue-it.net/festivalticketing/userdata/coach/wk1/coachW1w.css' type='text/css' />","javascripts":["//assets.queue-it.net/festivalticketing/userdata/coach/wk1/coachW1h.js","//assets.queue-it.net/festivalticketing/userdata/coach/wk1/coach.js"],"logoSrc":"//assets.queue-it.net/static/QueueFront/img/queue-it_logo_c0556c4b0e263943a08e2617c3550e37.png","toppanelIFrameSrc":null,"sidepanelIFrameSrc":null,"leftpanelIFrameSrc":null,"rightpanelIFrameSrc":null,"middlepanelIFrameSrc":null,"bottompanelIFrameSrc":null,"faviconUrl":"//assets.queue-it.net/festivalticketing/userdata/coach/wk1/favicon.ico","languages":[],"whatIsThisUrl":"https://queue-it.com/what-is-this/?customerId=festivalticketing&eventId=coachellawknd12018a&queueId=7065486b-36d2-4d20-a06b-4f6d53580286&language=en-US&hostUrl=aeg.festivalticketing.com","queueItLogoPointsToUrl":"https://queue-it.com?c=festivalticketing&e=coachellawknd12018a&q=7065486b-36d2-4d20-a06b-4f6d53580286","welcomeSoundUrls":["//assets.queue-it.net/static/QueueFront/css/sound/welcomeAudio_92a6592f5d4e6b14efdcc82656ba4273.mp3","//assets.queue-it.net/static/QueueFront/css/sound/welcomeAudio_0155ec254497c461d68f0de9114780ae.ogg","//assets.queue-it.net/static/QueueFront/css/sound/welcomeAudio_ccd6b626b658cceeab6f2c1ef926e694.wav"],"AppleTouchIconUrl":"//assets.queue-it.net/static/QueueFront/img/queue-it_appletouch_d72f69f877cca3aa3f8ddb61337a4714.png","DocumentTitle":"Coachella 2018 Queue"},"layout":{"languageSelectorVisible":false,"logoVisible":false,"dynamicMessageVisible":true,"reminderEmailVisible":false,"expectedServiceTimeVisible":false,"highLoadVisible":false,"queueNumberVisible":false,"usersInLineAheadOfYouVisible":false,"whichIsInVisible":false,"sidePanelVisible":false,"topPanelVisible":false,"leftPanelVisible":false,"rightPanelVisible":false,"middlePanelVisible":false,"bottomPanelVisible":false,"usersInQueueVisible":false,"queueIsPausedVisible":false,"reminderVisible":false,"servicedSoonVisible":false,"firstInLineVisible":false,"layoutVersion":1496272606,"queueNumberLoadingVisible":false,"progressVisible":true,"isRedirectPromptDialogEnabled":false,"isQueueitFooterHidden":false},"forecastStatus":"NotReadyYet","layoutName":"Coachella","layoutVersion":1496272606,"updateInterval":30000},
			        captchaVerifyEndpoint: 'https://eu-west-1-captchaverify-api.queue-it.net/captchaverify',
			        funCaptchaPublicKey: 'C08EDFEF-9E74-D179-FD28-3687C466FA05',
			        showCaptcha: false,
			        doGetStatusUpdate: true 
			      
			        });

			    window.queueViewModel.init();

			    $(document).ready(function() {
			        ko.applyBindings(window.queueViewModel);
			        window.queueViewModel.onReady();
			    });

			</script>
			    <script type="text/javascript" src="./Coachella 2018 Queue_files/coachW1h.js" crossorigin="anonymous"></script><script type="text/javascript" src="./Coachella 2018 Queue_files/coach.js" crossorigin="anonymous"></script>
			    
			    <style>
			    .before .queueElement, .queue .beforeElement {
			        display: none;
			    }
			</style>
			</head>
			<body data-pageid="queue" data-culture="en-US" class="queue">
			    <div id="noscript" style="display: none;">
			    <div id="BodyTop_nojavascriptenabled" class="nojavascript alert alert-error">
			        JavaScript is not enabled in your browser. Please enable JavaScript for automated updates. If you do not want to enable JavaScript, you must manually refresh this page to view your place in line. You can also click <a href="http://go.festivalticketing.com/?c=festivalticketing&amp;e=coachellawknd12018a&amp;q=7065486b-36d2-4d20-a06b-4f6d53580286&amp;cid=en-US&amp;l=Coachella&amp;javascriptdisabled=1">here</a> to automatically reload this page.
			    </div>
			</div>
			<script type="text/javascript">
			    document.getElementById("noscript").style.display = 'none';
			</script>

			    <div id="toppanel" style="display: none;" data-bind="visible: layout.topPanelVisible">
			    <iframe id="toppanel_iframe" src="./Coachella 2018 Queue_files/saved_resource.html" allowtransparency="true" frameborder="0" scrolling="no" data-bind="iframe-src-attr: { src: texts.toppanelIFrameSrc }"></iframe>
			</div>

			    <div id="wrapper"><div class="client-header"></div>
			        <div id="leftpanel" style="display: none;" data-bind="visible: layout.leftPanelVisible">
			    <iframe id="leftpanel_iframe" src="./Coachella 2018 Queue_files/saved_resource(1).html" allowtransparency="true" frameborder="0" scrolling="no" data-bind="iframe-src-attr: { src: texts.leftpanelIFrameSrc }"></iframe>
			</div>

			        <div id="main"><div class="block-logo"><span><img src="./Coachella 2018 Queue_files/logo.svg"></span></div><div class="date-display"><span class="w1">WEEKEND ONE</span><span class="w2">April 13-15, 2018</span></div>
			            <div id="main-top">
			    <div id="language-selector" title="select display language" style="display: none;" data-bind="visible: texts.languages().length &gt; 1 &amp;&amp; layout.languageSelectorVisible">
			        <select id="SharedTexts_Languages_SelectedValue" data-bind="options: texts.languages, optionsText:&#39;text&#39;, optionsValue:&#39;value&#39;, value:texts.selectedLanguage "></select>
			    </div>
			    <div id="main-top-message">
			        <p class="item">
			            <span id="lbCookieInfo"></span>
			        </p>
			    </div>
			</div>

			            <div id="main_t" class="t"></div>
			            <div id="main_c" class="c">
			                <div id="content">
			                    <div id="header">
			    <h1 class="logo" style="display: none;" data-bind="visible: layout.logoVisible">
			        <img src="./Coachella 2018 Queue_files/queue-it_logo_c0556c4b0e263943a08e2617c3550e37.png" id="imgCustomerLogo" alt="logo" longdesc="company logo for Gingerbread Shed Corporation" data-bind="attr: { src: texts.logoSrc }">
			    </h1>
			    <h2>
			        <span id="lbHeaderH2" data-bind="html: texts.header">You are now in line</span>
			    </h2>
			    <div id="headerparagraph">
			        <span id="lbHeaderP" data-bind="html: texts.body">You are now in line for the Coachella Weekend 1 advance sale. When it is your turn, you will have 10 minutes to enter the site and purchase your passes. If you do not enter the purchase page within that 10 minute period you will lose your place in line.</span>
			    </div>
			    <div id="whatisthis">
			        <a id="hlThisIsQueueIt" href="https://queue-it.com/what-is-this/?customerId=festivalticketing&amp;eventId=coachellawknd12018a&amp;queueId=7065486b-36d2-4d20-a06b-4f6d53580286&amp;language=en-US&amp;hostUrl=aeg.festivalticketing.com" target="_blank" data-bind="attr: { href: texts.whatIsThisUrl }">What is this?</a>
			    </div>
			</div>
			                    <div id="divConfirmRedirectModal" class="confirmredirectmodal queueElement" data-bind="visible: showConfirmRedirectDialog()" style="display: none">
			    <div id="divConfirmRedirectModal_Content">
			        <h2 id="h2ConfirmRedirect">Your turn started at</h2>
			        <h3 data-bind="{visible:ticket.windowStartTime()}" style="display: none;">
			            <span data-bind="{text:ticket.windowStartTime()}"></span>
			            <span data-bind="{text:ticket.timeZonePostfix()}">PST</span>
			        </h3>
			        <p id="pConfirmRedirect">Please confirm that you want to proceed to the site as soon as possible</p>
			       
			        <button id="buttonConfirmRedirect" type="button" class="btn" data-bind="click: setActiveClient"><span class="l">Yes, please</span><span class="r">&nbsp;</span></button>
			    </div>
			</div>

			                    
			                    <div class="warning-box">

			    <p class="beforeElement">
			        <span id="MainPart_lbEventStartsAtText">The onsale will begin at:</span>
			        <span id="MainPart_lbEventStartTime" data-bind="text: ticket.eventStartTimeFormatted">11:00 AM</span>
			        <span id="MainPart_lbEventStartTimeTimeZonePostfix" data-bind="text: ticket.timeZonePostfix">PST</span>
			    </p>

			    <div id="MainPart_divProgressbar" class="progressbar queueElement" data-bind="visible: layout.progressVisible">
			        <div id="MainPart_divProgressbar_Progress" class="progress" style="width: 18%;">
			            <div id="MainPart_divProgressbar_Progress_Runner" class="runner" data-bind="css: { paused: layout.queueIsPausedVisible }"></div>
			        </div>
			        <div id="MainPart_divProgressbar_Clear" class="clear"></div>
			    </div>
			</div>
			<div id="MainPart_lbManualUpdateWarning" style="display: none" class="alert alert-error" data-bind="visible: !isRunning()">You have lost your connection to the line. Please check your internet connection.</div>
			<div id="MainPart_divProgressbarBox" data-bind="visible: isRunning()" class="processbar-box">
			    <div id="MainPart_divProgressbarBox_Holder" class="holder">
			        <!-- before/idle -->
			        <div id="defaultCountdown" class="beforeElement hasCountdown"><div class="finished">You will now be redirected to the line. </div></div>
			        <p class="larger beforeElement">
			            </p><div id="MainPart_divUsersInQueueCount" style="display: none;" data-bind="visible: layout.usersInQueueVisible">
			                <span id="MainPart_lbUsersInQueueCountText">Total number of users in line:</span>
			                <span id="MainPart_lbUsersInQueueCount" data-bind="text: ticket.usersInQueue"></span>
			            </div>
			        <p></p>

			        <!-- queue -->
			        <p id="MainPart_pProgressbarBox_Holder_Larger" class="larger queueElement">
			            <span style="display: none;" data-bind="visible: layout.queueNumberVisible">
			                <span id="MainPart_lbQueueNumberText">Your number in line:</span>
			                <span id="MainPart_lbQueueNumber" data-bind="text: ticket.queueNumber"></span>
			                <br id="MainPart_lbQueueNumberLineBreak">
			            </span>
			            <span id="MainPart_lbUsersInLineAheadOfYouText" style="display: none;" data-bind="visible: layout.usersInLineAheadOfYouVisible">Number of users in line ahead of you:</span>
			            <span id="MainPart_lbUsersInLineAheadOfYou" style="display: none;" data-bind="visible: layout.usersInLineAheadOfYouVisible, text: ticket.usersInLineAheadOfYou"></span>
			            <span id="first-in-line" style="display: none;" data-bind="visible: layout.firstInLineVisible">It is your turn</span>
			            <br id="MainPart_lbUsersInLineAheadOfYouLineBreak">
			            <span id="MainPart_lbExpectedServiceTimeText" style="display: none;" data-bind="visible: layout.expectedServiceTimeVisible">Expected arrival time on the website:</span>
			            <span id="MainPart_lbExpectedServiceTime" style="display: none;" data-bind="visible: layout.expectedServiceTimeVisible, text: ticket.expectedServiceTime"></span>
			            <span id="MainPart_lbExpectedServiceTimeTimeZonePostfix" style="display: none;" data-bind="visible: layout.expectedServiceTimeVisible, text: ticket.timeZonePostfix">PST</span>
			            <span id="queue-paused" style="display: none;" data-bind="visible: layout.queueIsPausedVisible">The line is paused.</span>
			            <br id="MainPart_lbExpectedServiceTimeLineBreak">
			            <span id="serviced-soon" style="display: none;" data-bind="visible: layout.servicedSoonVisible()">
			                <span data-bind="visible: !layout.servicedSoonDelayVisible()">Thank you for waiting. You are now being redirected to the website.</span>
			                <span data-bind="visible: layout.servicedSoonDelayVisible()" style="display: none;">You're almost there. You will be redirected to the website as soon as possible.</span>
			            </span>
			            <span id="MainPart_lbWhichIsInText" style="display: none;" data-bind="visible: layout.whichIsInVisible">Your estimated wait time is:</span>
			            <span id="MainPart_lbWhichIsIn" style="display: none;" data-bind="visible: layout.whichIsInVisible, text: ticket.whichIsIn"></span>
			        </p>
			        <ul id="MainPart_ulProgressbarBox_Holder_Processbar" class="processbar" style="width:21px">
			            <li id="defaultViewPb1" class=""><span>&nbsp;</span></li>
			        </ul>
			        <div id="MainPart_divProgressbarBox_Holder_LastUpdateTime" data-bind="visible: isRunning">
			            <span id="MainPart_lbLastUpdateTime" title="Everything is fine, as long the &#39;Status last updated&#39; time is updated regularly.">Status last updated:</span>
			            <span id="MainPart_lbLastUpdateTimeText" data-bind="text: ticket.lastUpdated">11:32:56 AM</span>
			            <span id="MainPart_lbLastUpdateTimeTextTimeZonePostfix" data-bind="text: ticket.timeZonePostfix">PST</span>
			        </div>
			    </div>
			</div>
			<div id="MainPart_divBlock" class="block beforeElement">
			    <p></p>
			</div>
			                    <div id="MainPart_divTimeBox" class="time-box" style="display: none;" data-bind="visible: layout.dynamicMessageVisible() &amp;&amp; message()">
			    <div id="MainPart_divTimeBox_Holder" class="holder" data-bind="if: message"></div>
			</div>
			                    <div id="MainPart_frmReminder2" name="Notify me" class="reminder-form queueElement" style="display: none;" data-bind="visible: layout.reminderVisible">
			    <h2><span id="MainPart_lbNotyfyMeText">Please notify me when it is my turn:</span></h2>
			    <div id="MainPart_divEmailInput" class="row" style="display: none;" data-bind="visible: layout.reminderEmailVisible">
			        <form id="MainPart_FormEmailInput" name="Email notification form">
			            <div style="display: none;">Enter email address</div>
			            <input name="ctl00$MainPart$inpEmailAddress" type="text" id="MainPart_inpEmailAddress" title="Enter email address" class="item-input" placeholder="Enter email address" data-bind="value: emailAddress">
			            <button type="submit" class="btn" id="aUpdateEmail" data-bind="click: updateNotify"><span class="l">Notify me by e-mail</span><span class="r">&nbsp;</span></button>
			        </form>
			        <div id="divEmailStatusFrame">
			            <div id="divEmailStatusImage"></div><div id="divEmailStatus"></div>
			        </div>
			    </div>
			</div>
			                        <div id="MainPart_divExitLine" class="block queueElement">
			        <p><a href="http://go.festivalticketing.com/exitline.aspx?c=festivalticketing&amp;e=coachellawknd12018a&amp;q=7065486b-36d2-4d20-a06b-4f6d53580286&amp;cid=en-US&amp;l=Coachella" id="MainPart_aExitLine">Leave the line</a> (You will lose your place)</p>
			    </div>

			                </div>
			                <div id="middlepanel" style="display: none;" data-bind="visible: layout.middlePanelVisible">
			    <iframe id="middlepanel_iframe" src="./Coachella 2018 Queue_files/saved_resource(2).html" allowtransparency="true" frameborder="0" scrolling="no" data-bind="iframe-src-attr: { src: texts.middlepanelIFrameSrc }"></iframe>
			</div>

			                <div id="footer">
			    <div id="footer-disclaimer" data-bind="html: texts.disclaimerText"></div>


			        <div id="footer-direct-link" style="" data-bind="visible: hasQueueId">Queue ID: <a href="http://go.festivalticketing.com/?c=festivalticketing&amp;e=coachellawknd12018a&amp;q=7065486b-36d2-4d20-a06b-4f6d53580286&amp;cid=en-US&amp;l=Coachella" id="hlLinkToQueueTicket2" data-bind="text: queueId">7065486b-36d2-4d20-a06b-4f6d53580286</a></div>
			            <div id="footer-queueit-link">
			            <strong class="by">
			                <a href="https://queue-it.com/?c=festivalticketing&amp;e=coachellawknd12018a&amp;q=7065486b-36d2-4d20-a06b-4f6d53580286" id="aThisIsQueueIt" target="_blank">QUEUE-IT</a>
			            </strong>
			        </div>
			</div>
			            </div>
			            <div id="MainPart_divVersionInfo" class="b" title="4.0.196.0"></div>
			        </div>
			        <div id="rightpanel" style="display: none;" data-bind="visible: layout.rightPanelVisible">
			    <iframe id="rightpanel_iframe" src="./Coachella 2018 Queue_files/saved_resource(3).html" allowtransparency="true" frameborder="0" scrolling="no" data-bind="iframe-src-attr: { src: texts.rightpanelIFrameSrc }"></iframe>
			</div>

			    </div>
			    <div id="sidebar" style="display: none;" data-bind="visible: layout.sidePanelVisible">
			    <iframe id="sidebar_iframe" src="./Coachella 2018 Queue_files/saved_resource(4).html" allowtransparency="true" frameborder="0" scrolling="no" data-bind="iframe-src-attr: { src: texts.sidepanelIFrameSrc }"></iframe>
			</div>

			    <div id="bottompanel" style="display: none;" data-bind="visible: layout.bottomPanelVisible">
			    <iframe id="bottompanel_iframe" src="./Coachella 2018 Queue_files/saved_resource(5).html" allowtransparency="true" frameborder="0" scrolling="no" data-bind="iframe-src-attr: { src: texts.bottompanelIFrameSrc }"></iframe>
			</div>





			<script type="text/javascript" id="">(function(a,e,f,g,b,c,d){a[b]||(a.GlobalSnowplowNamespace=a.GlobalSnowplowNamespace||[],a.GlobalSnowplowNamespace.push(b),a[b]=function(){(a[b].q=a[b].q||[]).push(arguments)},a[b].q=a[b].q||[],c=e.createElement(f),d=e.getElementsByTagName(f)[0],c.async=1,c.src=g,d.parentNode.insertBefore(c,d))})(window,document,"script","//tracking.aegpresents.com/universalpixel/cf.js","cf2");
			window.cf2("newTracker","cf","//b.aegpresents.com",{appId:document.location.hostname,respectDoNotTrack:!0,contexts:{webPage:!0,performanceTiming:!0,gaCookies:!0,geolocation:!1,augurIdentityLite:!0,optimizelyExperiments:!0,optimizelyStates:!0,optimizelyVariations:!0,optimizelyVisitor:!0,optimizelyAudiences:!0,optimizelyDimensions:!0,optimizelySummary:!0}});window.cf2("enableActivityTracking",30,10);window.cf2("trackPageView");window.cf2("enableLinkClickTracking");</script><script type="text/javascript" id="">var cfe={ticketer:"gingerbreadshed",event:"queue"};cfe.eventName=document.title;dataLayer.push(cfe);</script>


			<script type="text/javascript" id="">!function(b,e,f,g,a,c,d){b.fbq||(a=b.fbq=function(){a.callMethod?a.callMethod.apply(a,arguments):a.queue.push(arguments)},b._fbq||(b._fbq=a),a.push=a,a.loaded=!0,a.version="2.0",a.queue=[],c=e.createElement(f),c.async=!0,c.src=g,d=e.getElementsByTagName(f)[0],d.parentNode.insertBefore(c,d))}(window,document,"script","//connect.facebook.net/en_US/fbevents.js");fbq("init","217011611820041");fbq("track","PageView");
			fbq("track","ViewContent",{host:location.hostname,genre:google_tag_manager["GTM-M9NK8HW"].macro('gtm4'),subGenre:google_tag_manager["GTM-M9NK8HW"].macro('gtm5'),artistName:google_tag_manager["GTM-M9NK8HW"].macro('gtm6'),artistId:google_tag_manager["GTM-M9NK8HW"].macro('gtm7'),eventName:google_tag_manager["GTM-M9NK8HW"].macro('gtm8')});</script>

			<script type="text/javascript" id="">!function(b,e,f,g,a,c,d){b.fbq||(a=b.fbq=function(){a.callMethod?a.callMethod.apply(a,arguments):a.queue.push(arguments)},b._fbq||(b._fbq=a),a.push=a,a.loaded=!0,a.version="2.0",a.queue=[],c=e.createElement(f),c.async=!0,c.src=g,d=e.getElementsByTagName(f)[0],d.parentNode.insertBefore(c,d))}(window,document,"script","//connect.facebook.net/en_US/fbevents.js");fbq("init","291582177684507");fbq("track","PageView");
			fbq("track","ViewContent",{host:location.hostname,eventName:google_tag_manager["GTM-M9NK8HW"].macro('gtm9'),genre:google_tag_manager["GTM-M9NK8HW"].macro('gtm10'),subGenre:google_tag_manager["GTM-M9NK8HW"].macro('gtm11'),artistName:google_tag_manager["GTM-M9NK8HW"].macro('gtm12'),artistId:google_tag_manager["GTM-M9NK8HW"].macro('gtm13'),venueName:google_tag_manager["GTM-M9NK8HW"].macro('gtm14'),venueCity:google_tag_manager["GTM-M9NK8HW"].macro('gtm15'),venueZip:google_tag_manager["GTM-M9NK8HW"].macro('gtm16')});</script>
			<script type="text/javascript" id="">!function(b,e,f,g,a,c,d){b.fbq||(a=b.fbq=function(){a.callMethod?a.callMethod.apply(a,arguments):a.queue.push(arguments)},b._fbq||(b._fbq=a),a.push=a,a.loaded=!0,a.version="2.0",a.queue=[],c=e.createElement(f),c.async=!0,c.src=g,d=e.getElementsByTagName(f)[0],d.parentNode.insertBefore(c,d))}(window,document,"script","//connect.facebook.net/en_US/fbevents.js");fbq("init","217011611820041");fbq("track","PageView");
			fbq("track","ViewContent",{host:document.location.hostname,ticketer:google_tag_manager["GTM-M9NK8HW"].macro('gtm17'),eventName:google_tag_manager["GTM-M9NK8HW"].macro('gtm18'),genre:google_tag_manager["GTM-M9NK8HW"].macro('gtm19'),subGenre:google_tag_manager["GTM-M9NK8HW"].macro('gtm20'),artistName:google_tag_manager["GTM-M9NK8HW"].macro('gtm21'),artistId:google_tag_manager["GTM-M9NK8HW"].macro('gtm22'),venueName:google_tag_manager["GTM-M9NK8HW"].macro('gtm23'),venueCity:google_tag_manager["GTM-M9NK8HW"].macro('gtm24'),venueZip:google_tag_manager["GTM-M9NK8HW"].macro('gtm25')});</script>

			</body></html>`,
			expected: "http://go.festivalticketing.com/?c=festivalticketing&e=coachellawknd12018a&q=7065486b-36d2-4d20-a06b-4f6d53580286&cid=en-US&l=Coachella",
		},
	}

	for msg, tc := range testcases {
		actual, err := parseQueueLink(tc.html)
		require.NoError(t, err, msg)
		assert.Equal(t, tc.expected, actual, msg)
	}
}
