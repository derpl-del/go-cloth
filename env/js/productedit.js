$(document).ready(function () {
    var pid = getUrlVars()["product_id"];
    url = "/api/v1/productdetail?product_id=" + pid;
    function getUrlVars() {
        var vars = {};
        var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function (m, key, value) {
            vars[key] = value;
        });
        return vars;
    }

    $.get(url, function (data) {
        var response = jQuery.parseJSON(data);
        imgloc = "../" + response.imgsource;
        $("#img").attr("src", imgloc);
        $("#product-name").text(response.productname);
        $("#product-code").text(response.productcode);
        $("#product-price").text(response.productprice);
        $("#product-gender").text(response.productgender);
        $("#product-size").text(response.productsize);
        $("#product-category").text(response.productcategory);
        $("#product-amount").text(response.productamount);
    });

    $("#edit-btn").click(function () {
        var idr = $("#product-price").text();
        $("#edit-modal").modal('show');
        $('#edit-name').val($("#product-name").text());
        $('#edit-price').val(idr.substring(4, idr.length));
        $('#edit-amount').val($("#product-amount").text());
        $('#edit-size').val($("#product-size").text());
        $('#edit-category').val($("#product-category").text());
        $('#edit-gender').val($("#product-gender").text());
    });

    $("#delete-btn").click(function () {
        $("#delete-modal").modal('show');
    });

    $("#edit-confirm").click(function () {
        urldelete = "/api/v1/editproduct";
        var req = JSON.stringify({ "productname": $("#edit-name").val(), "productcode": pid, "productprice": $('#edit-price').val(), "productgender": $('#edit-gender').val(), "productsize": $('#edit-size').val(), "productcategory": $('#edit-category').val(), "productamount": $('#edit-amount').val() });
        $.post(urldelete, req, function (resdata, status) {
            var res2 = jQuery.parseJSON(resdata);
            if (res2.ErrorCode == "0000") {
                sessionStorage.setItem("result", "0000");
                window.location.replace("/shopkeeper/product_list");
            } else {
                $("#preview-modal").modal('hide');
            }
        });
    });


    $("#delete-confirm").click(function () {
        urldelete = "/api/v1/deleteproduct?product_id=" + pid;
        $.get(urldelete, function (data) {
            var res2 = jQuery.parseJSON(data);
            if (res2.ErrorCode == "0000") {
                sessionStorage.setItem("result", "0000");
                window.location.replace("/shopkeeper/product_list");
            } else {
                $("#delete-modal").modal('hide');
            }
        });
    });

});