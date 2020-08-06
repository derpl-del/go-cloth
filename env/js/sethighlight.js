var opt1 = '<option value="';
var opt2 = '">';
var opt3 = '</option>';
var response;

$(document).ready(function () {
    url = "/api/v1/getproductall?";
    $.get(url, function (data) {
        response = jQuery.parseJSON(data);
        $.each(response.ListProduct, function (i, item) {
            ReturnOption(item.productcode, item.productname);
        });
    });

    function ReturnOption(code, value) {
        var build = opt1 + code + opt2 + value + opt3;
        $("#set-1").append(build);
        $("#set-2").append(build);
        $("#set-3").append(build);
    };



    $(".form-control").change(function () {
        var value = $(this);
        var prevVal = value.data("prev");
        var otherSelects = $(".form-control").not(this);
        if ($(this).val() == "none") {

        } else {
            otherSelects.find('option[value="' + $(this).val() + '"]').attr('disabled', true);
        }
        if (prevVal) {
            otherSelects.find('option[value="' + prevVal + '"]').attr('disabled', false);
        }
        value.data("prev", value.val());
    });

    $("#confirm-btn").click(function () {
        $("#preview-modal").modal("show");
        $(".carousel-inner").find("div").attr('pid', "");
        $(".carousel-inner").find("div").find("img").attr('src', "");
        $.each(response.ListProduct, function (i, item) {
            var img = "../" + item.imgsource
            if (item.productcode == $("#set-1").val()) {
                $("#img1").find("img").attr('src', img);
                $("#img1").attr('pid', item.productcode);
            } else if (item.productcode == $("#set-2").val()) {
                imgloc = "../" + item.imgsource;
                $("#img2").find("img").attr('src', img);
                $("#img2").attr('pid', item.productcode);
            } else if (item.productcode == $("#set-3").val()) {
                imgloc = "../" + item.imgsource;
                $("#img3").find("img").attr('src', img);
                $("#img3").attr('pid', item.productcode);
            } else { }
        });
    });

    $("#apply-set").click(function () {
        var ListData = {}
        var product = []
        $(".carousel-inner").find("div").each(function (i) {
            var productcode = $(this).attr("pid");
            if (productcode != "") {
                var src = $(this).find("img").attr("src");
                product.push({ productcode: productcode, imgloc: src });
            }
            else {
            }
        });
        ListData['ListProduct'] = product;
        var reqJS = JSON.stringify(ListData);
        console.log(reqJS);
        url = "/api/v1/makehighlight";
        $.post(url, reqJS, function (resdata, status) {
            var res2 = jQuery.parseJSON(resdata);
            if (res2.ErrorCode == "0000") {
                window.location.replace("/");
            } else {
                window.location.replace("/404");
            }
        });
    });
});