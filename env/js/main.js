var div1 = '<div pid="';
var div2 = '" class="carousel-item ';
var div3 = ' text-center"><img src="';
var div4 = '" class="img-fluid" alt="Responsive image"></div>';

$(document).ready(function () {
    $.get("/api/v1/gethighlight", function (data) {
        var res2 = jQuery.parseJSON(data);
        $.each(res2.ListProduct, function (i, item) {
            ReturnComponent(i, item.productcode, item.imgloc);
        });
    });
    function ReturnComponent(i, code, img) {
        if (i == 0) { value = "active"; }
        else { value = "inactive"; }
        var build = div1 + code + div2 + value + div3 + img + div4;
        $(".carousel-inner").append(build);
    };
});