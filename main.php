<?php
$headers = array_filter($_SERVER, function ($header) {
    return in_array($header, ['CONTENT_TYPE', 'CONTENT_LENGTH'], true) || strpos($header, 'HTTP_') === 0;
}, ARRAY_FILTER_USE_KEY);

$headers = array_combine(
    array_map(function ($header) {
        if (strpos($header, 'HTTP_') === 0) {
            $header = substr($header, 5);
        }

        return str_replace(' ', '-', ucwords(str_replace('_', ' ', strtolower($header))));
    }, array_keys($headers)),
    array_values($headers)
);

$response['headers'] = $headers;
$response['method']  = $_SERVER['REQUEST_METHOD'];

$body = file_get_contents('php://input');

$response['body']  = (isset($_SERVER['CONTENT_TYPE']) && $_SERVER['CONTENT_TYPE'] === 'application/json') ? json_decode($body) : $body;
$response['query'] = $_GET;

header('Content-Type: application/json');
echo json_encode($response);
