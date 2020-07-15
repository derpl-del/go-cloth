var isname = false;
var isprice = false;
var isamount = false;
var issize = false;
var isgender = false;
var iscategory = false;
var isfile = false;
var msg_error;

$(document).ready(function () {
    $("#name").val("");
    $("#price").val("");
    $("#amount").val("");
    $("#size").val("");
    $("#gender").val("");
    $("#category").val("");
    $("#file").val("");

    $.get("/api/v1/generateproductid", function (data) {
        var response = jQuery.parseJSON(data);
        if (response.Location != "") {
            $("#form-id").attr("token", response.Location);
        } else {
            alert("Please Refresh Your Browser")
        }
    });

    $("#name").change(function () {
        var name = $("#name").val();
        if (name != "") {
            isname = true;
        } else {
            isname = false;
        }
    });

    $("#price").change(function () {
        var price = $("#price").val();
        if (price != "") {
            isprice = true;
        } else {
            isprice = false;
        }
    });

    $("#amount").change(function () {
        var amount = $("#amount").val();
        if (amount != "") {
            isamount = true;
        } else {

            isamount = false;
        }
    });

    $("#size").change(function () {
        var size = $("#size").val();
        if (size != "") {
            issize = true;
        } else {
            issize = false;
        }
    });

    $("#gender").change(function () {
        var gender = $("#gender").val();
        if (gender != "") {
            isgender = true;
        } else {
            isgender = false;
        }
    });

    $("#category").change(function () {
        var category = $("#category").val();
        if (category != "") {
            iscategory = true;
        } else {
            iscategory = false;
        }
    });

    $("#file").change(function () {
        var filedata = $("#file").val();
        if (filedata != "") {
            isfile = true;
        } else {
            isfile = false;
        }
    });

    $("#but_upload").click(function () {
        $(".form-group p").empty();
        msg_error = "invalid input";
        if (isname == true && isprice == true && isamount == true &&
            issize == true &&
            isgender == true &&
            iscategory == true &&
            isfile == true) {
            var fd = new FormData();
            var files = $('#file')[0].files[0];
            var ticket = $("#form-id").attr("token");
            fd.append('file', files);
            fd.append('token', ticket);
            $.ajax({
                url: '/api/v1/preview',
                type: 'post',
                data: fd,
                contentType: false,
                processData: false,
                success: function (data) {
                    var response = jQuery.parseJSON(data);
                    if (response.ErrorCode == "0000") {
                        imgloc = "../" + response.Location;
                        $("#img").attr("src", imgloc);
                        var productname = $("#name").val();
                        var productid = "Product Code : " + response.Productid;
                        var productprice = $("#price").val();
                        var pricep = "IDR " + productprice;
                        var productgender = $("#gender").val();
                        var productsize = $("#size").val();
                        var productcategory = $("#category").val();
                        var productamount = $("#amount").val();
                        $("#product-name").text(productname);
                        $("#product-code").text(productid);
                        $("#product-price").text(pricep);
                        $("#product-gender").text(productgender);
                        $("#product-size").text(productsize);
                        $("#product-category").text(productcategory);
                        $("#product-amount").text(productamount);
                        $("#preview").show(); // Display image element
                        $("#preview-modal").modal('show');
                    } else {
                        alert('file not uploaded');
                    }
                },
            });
        } else if (isname == false) {
            var top = $("#name").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(0) p").append(document.createTextNode(msg_error));
        } else if (isprice == false) {

            var top = $("#price").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(1) p").append(document.createTextNode(msg_error));
        } else if (isamount == false) {

            var top = $("#amount").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(2) p").append(document.createTextNode(msg_error));
        } else if (issize == false) {
            var top = $("#size").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(3) p").append(document.createTextNode(msg_error));
        } else if (isgender == false) {
            var top = $("#gender").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(4) p").append(document.createTextNode(msg_error));
        } else if (iscategory == false) {
            var top = $("#category").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(5) p").append(document.createTextNode(msg_error));
        } else if (isfile == false) {
            var top = $("#file").offset().top - $(".form-group").height();
            $('html, body').stop().animate({
                scrollTop: top
            }, 500);
            $(".form-group:eq(6) p").append(document.createTextNode(msg_error));
        }
        else { }
    });

    $("#create-product").click(function () {
        var productname = $("#product-name").text();
        var producttoken = $("#form-id").attr("token");
        var productcode = $("#product-code").text();
        var productprice = $("#product-price").text();
        var productgender = $("#product-gender").text();
        var productsize = $("#product-size").text();
        var productcategory = $("#product-category").text();
        var productamount = $("#product-amount").text();
        var imgsource = $("#preview img").attr("src");
        var req = JSON.stringify({ "productname": productname, "productcode": productcode, "productprice": productprice, "productgender": productgender, "productsize": productsize, "productcategory": productcategory, "productamount": productamount, "imgsource": imgsource, "producttoken": producttoken });
        $.post('/api/v1/productadd', req, function (resdata, status) {
            var res2 = jQuery.parseJSON(resdata);
            if (res2.ErrorCode == "0000") {
                window.location.replace("/shopkeeper");
            } else {
                $("#preview-modal").modal('hide');
            }
        });
    });
});