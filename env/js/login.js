var $signdiv1 = "<div id='notif-alert' class='alert alert-warning alert-dismissible show' role='alert'>"
var $signdiv2 = "<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>"
var $succdiv1 = "<div id='notif-alert' class='alert alert-success alert-dismissible show' style='opacity: 1;' role='alert'>"
$(document).ready(function () {
    var ErrorCode = sessionStorage.getItem("ErrorCode");
    if (ErrorCode == "0000") {
        greeting = "Success Created Account";
        $("#notif-result").append($succdiv1 + greeting + $signdiv2)
        sessionStorage.setItem("ErrorCode", "");
    }
    else {
        $('#notif-alert').alert('close')
    }
    $.get("/GenerateToken", function (data) {
        var response = jQuery.parseJSON(data);
        if (response.ErrorMessage != "") {
            $("#logo-jumbo").attr("token", response.ErrorMessage);
        } else {
            alert("Please Refresh Your Browser")
        }
    });
    $('#LogInPage').on('hidden.bs.modal', function (e) {
        $("form").trigger("reset");
    })
    $('#SignUpPage').on('hidden.bs.modal', function (e) {
        $("form").trigger("reset");
    })

    $("#loginform").keypress(function (e) {
        if (e.which == 13) {
            $("#login_send").trigger("click");
        }
    });

    $("#signupform").keypress(function (e) {
        if (e.which == 13) {
            $("#signup_send").trigger("click");
        }
    });

    $("#login_send").click(function () {
        $('#notif-alert').alert('close')
        var username = $("#UsernameIn").val();
        var password = $("#PasswordIn").val();
        var ticket = $("#logo-jumbo").attr("token");
        var data = JSON.stringify({ "username": username, "password": password, "ticket": ticket });
        $.post("/SignInUser", data, function (data, status) {
            var response = jQuery.parseJSON(data);
            if (response.ErrorCode == "0000") {
                window.location.replace("/");
            } else {
                greeting = response.ErrorMessage;
                $("#notif-in-success").append($signdiv1 + greeting + $signdiv2);
                if (response.ErrorCode == "0001") {
                    $("form").trigger("reset");
                }
                else if (response.ErrorCode == "0002") {
                    $("#PasswordIn").val("");
                }
            }
        });
    });

    $("#signup_send").click(function () {
        $('#notif-alert').alert('close')
        var username = $("#UsernameUp").val();
        var password = $("#PasswordUp").val();
        var data = JSON.stringify({ "username": username, "password": password });
        $.post("/SignUpUser", data, function (data, status) {
            var response = jQuery.parseJSON(data);
            if (response.ErrorCode == "0000") {
                sessionStorage.setItem("ErrorCode", "0000");
                document.location.reload();
            } else {
                greeting = response.ErrorMessage;
                $("#notif-up-success").append($signdiv1 + greeting + $signdiv2);
                if (response.ErrorCode == "1001" || (response.ErrorCode == "1003")) {
                    $("#UsernameUp").val("");
                }
                else if ((response.ErrorCode == "1004")) {
                    $("#PasswordUp").val("");
                } else {
                    $("form").trigger("reset");
                }
            }
        });
    });
});