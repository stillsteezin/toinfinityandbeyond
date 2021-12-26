<?php
session_start();

$con = mysqli_connect('localhost', 'root', ' ', 'phplogin');
mysqli_select_db($con, 'phplogin');
$id = $_POST['user']
$pass = $_POST['password'];	

$s = " select * from usertable where name = '$username' && password = '$password'";

$result = mysqli_query($con, $s);

$num = mysqli_num_rows($result);

if($num == 1){
	header('location:trackplants.html');
}else{
	echo 'Incorrect username and/or password!';
	header('location:login.php');
}

if ( mysqli_connect_errno() ) {
	// If there is an error with the connection, stop the script and display the error.
	exit('Failed to connect to MySQL: ' . mysqli_connect_error());
}



$stmt->close();

?>
