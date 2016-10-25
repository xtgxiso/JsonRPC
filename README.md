# JsonRPC
GO语言json协议的rpc服务和调用
php调用go的服务

<?php
$client = new JsonRPC("127.0.0.1", 1215);
$r = $client->Call("MyMath.Add",array('num1'=>1,'num2'=>2));
var_export($r);
