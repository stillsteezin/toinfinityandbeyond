<?php

// Update the path below to your autoload.php,
// see https://getcomposer.org/doc/01-basic-usage.md
require_once '/path/to/vendor/autoload.php';

use Twilio\Rest\Client;

// Find your Account SID and Auth Token at twilio.com/console
// and set the environment variables. See http://twil.io/secure
$sid = getenv("ACb559af4eb858c653d7fb9fbab63e8909");
$token = getenv("TWILIO_AUTH_TOKEN");
$twilio = new Client($sid, $token);

$message = $twilio->messages
                  ->create("+15558675310", // to
                           [
                               "body" => "Your plants are ready to harvest!",
                               "from" => "+16203191314",
                               "statusCallback" => "http://postb.in/1234abcd"
                           ]
                  );

print($message->sid);
