<?php




$host = "localhost";
$username = "root";
$password = "";

$conn = mysqli_connect($host, $username, $password);

if (!$conn) {
    echo "<p>failed to connect db</p>";
    die("Exit<br/>");
}
if ($_SERVER['REQUEST_METHOD'] === 'POST') {
    $createDbNTable = "CREATE DATABASE formData;";
    $createDbNTable .= "CREATE TABLE formData.users (
    id INT(6) UNSIGNED AUTO_INCREMENT NOT NULL PRIMARY KEY,
    name VARCHAR(30) NOT NULL,
    email VARCHAR(100) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    password VARCHAR(100) NOT NULL);";

    try {
        if (mysqli_multi_query($conn, $createDbNTable)) {
            echo "<p>DB and table created successfully <p>";
            // echo "Error " . mysqli_error($conn);
        } else {
            echo "Error: " . mysqli_connect_error();
        }
    } catch (Exception $e) {
        echo "<small>Database or table may already exists</small>";
    }


    $name = $_POST['name'];
    $email = $_POST['email'];
    $phone = $_POST['phone'];
    $pswd = $_POST['pswd'];

    $sql = "INSERT INTO formData.users (name, email, phone, password) 
            VALUES ('$name', '$email', '$phone', '$pswd')
            ";
    mysqli_query($conn, $sql);
    $recent_record_id = mysqli_insert_id($conn);
    echo "<p>data created successfully. - $recent_record_id</p>";
}

mysqli_close($conn);
echo "<p>Connection close</p>";
?>