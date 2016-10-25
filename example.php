<?php
include("JsonRPC.php");
$client = new JsonRPC("127.0.0.1", 1215);
$r = $client->Call("MyMath.Add",array('num1'=>1,'num2'=>2));
var_export($r);
echo "<br/>";
$r = $client->Call("MyMath.Sub",array('num1'=>'1','num2'=>'2'));
var_dump($r);
