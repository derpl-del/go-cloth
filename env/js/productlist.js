var page = 1;
var div1 = "<div class='col-md-4'><div class='card mb-4 shadow-sm' style='width: 18rem;'><img src='";
var div2 = "' class='card-img-top' alt='...'> <div class='card-body'> <h4 id='product-name'>";
var div3 = "</h4> <p id='product-id'>";
var div4 = "</p> <a type='button' href='"
var div5 = "' class='btn btn-sm btn-outline-secondary'>Edit</a> </div> </div></div>";
var totalpage;
var page1 = "<li class='page-item'><a class='page-link' href='?page=";
var page2 = "'>";
var page3 = "</a></li>";
var cpage1 = "<li class='page-item active' aria-current='page'> <a class='page-link' href='#'>";
var cpage2 = " <span class='sr-only'>(current)</span></a> </li>";
var $succdiv1 = "<div id='notif-login' class='alert alert-success alert-dismissible show' role='alert'>"
var $newdiv2 = "<button type='button' class='close' data-dismiss='alert' aria-label='Close'><span aria-hidden='true'>&times;</span></button></div>"

$(document).ready(function () {
    var page = getUrlVars()["page"];
    if (page === undefined) { page = 1; }
    url = "/api/v1/productall?page=" + page;
    var result = sessionStorage.getItem("result");
    if (result == "0000") {
        greeting = "Success";
        $("#notif-result").append($succdiv1 + greeting + $newdiv2);
        sessionStorage.setItem("result", "");
    }
    else {
    }
    $.get(url, function (data) {
        var response = jQuery.parseJSON(data);
        totalpage = response.Totalpage;
        $.each(response.ListProduct, function (i, item) {
            ReturnCard(item.imgsource, item.productname, item.productcode);
        });
        CreatePagenation(totalpage, page);
    });

    function ReturnCard(src, tittle, code) {
        var img = "../" + src
        var link = "/shopkeeper/product_edit?product_id=" + code
        var datadiv = div1 + img + div2 + tittle + div3 + code + div4 + link + div5;
        $("#data-product").append(datadiv);
    }

    function CreatePagenation(total, current) {
        var pagestr = ""
        for (i = 1; i <= total; i++) {
            if (i == current) {
                pagestr = pagestr + cpage1 + current + cpage2
            } else {
                pagestr = pagestr + page1 + i + page2 + i + page3;
            }
        }
        $("#page-list").append(pagestr);
    }

    function getUrlVars() {
        var vars = {};
        var parts = window.location.href.replace(/[?&]+([^=&]+)=([^&]*)/gi, function (m, key, value) {
            vars[key] = value;
        });
        return vars;
    }

});