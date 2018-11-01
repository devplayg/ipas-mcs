/**
 *
 */

var StartupColor = "#e1e5ec", // blue-chambray
    ShockColor = "#3598DC", // blue
    SpeedingColor = "#f4902f", // warning(yellow) // #f4902f
    ProximityColor = "#E7505A"; // red

// Notification
toastr.options = {
    "closeButton": true,
    "debug": false,
    "newestOnTop": false,
    "progressBar": true,
    "positionClass": "toast-top-right",
    "preventDuplicates": true,
    "onclick": null,
    "showDuration": "300",
    "hideDuration": "1000",
    "timeOut": "20000",
    "extendedTimeOut": "10000",
    "showEasing": "swing",
    "hideEasing": "linear",
    "showMethod": "fadeIn",
    "hideMethod": "fadeOut"
}

updateNews();
var timer = setInterval(function() {
    try {
        updateNews();
    } catch(err) {
        console.log( err );
    }
}, 600 * 1000); // Demo



$.ajaxSetup({ cache:false });
var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === "object" ) {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $( "meta[name=_xsrf]" ).attr( "content" );
        var headers = options.headers || {};
        var domain = document.domain.replace(/\./ig, '\\.');
        if (!/^(http:|https:).*/.test(url) || eval('/^(http:|https:)\\/\\/(.+\\.)*' + domain + '.*/').test(url)) {
            headers = $.extend(headers, {'X-Xsrftoken':xsrftoken});
        }
        options.headers = headers;
        return ajax(url, options);
    }
});

// jQuery
jQuery.fn.addHidden = function (name, value) {
    return this.each(function () {
        var input = $( "<input>" ).attr( "type", "hidden" ).attr( "name", name ).val( value );
        $( this ).append( $( input ) );
    });
};

// btn-global-message

$( document ).on('click', '.toast', function(){
    var a = $( this ).find( "a:first" ),
        messageId = a.data( "message-id" );

    if ( messageId > 0 ) { // 메시지 1건에 대한 마킹 처리 시
        $.get( "/message/gotit/" + messageId );
    } else { // Clear all toastr
        $.get( "/message/markAll" );
        toastr.clear();
    }
});

// Mask
$( ".mask-yyyymmddhhii" ).mask( "0000-00-00 00:00" );
$( ".mask-ipv4-cidr" ).mask( "099.099.099.099/09" );
$( ".mask-09999" ).mask( "09999" );
$( ".mask-0999" ).mask( "0999" );
$( ".mask-099" ).mask( "099" );

// Language
$( ".lang-changed" ).click( function() {
    var $e = $( this );
    var lang = $e.data( "lang" );
    $.cookie( "lang", lang, { path: "/", expires: 365} );
    window.location.reload();
});

// Sidebar
Layout.setSidebarMenuActiveLink( "match" )

// Selected menu
updateNavText();

//

/**
 * Functions
 */

function updateNavText() {
    var menu = $( ".page-sidebar-menu" ),
        el = null,
        url = location.pathname.toLowerCase();

    menu.find("li > a").each(function() {
        var path = $( this ).attr( "href" ).toLowerCase();
        if (path.length > 1 && url.substr( 1, path.length - 1 ) == path.substr( 1 )) {
            el = $( this );
            return;
        }
    });
    if ( el === null ) {
        return;
    }
    var parent = el.closest( "ul" ).parent().find( "a:first" ),
        depth2 = $.trim( el.text() ),
        depth1 = $.trim( parent.text() );

    $( ".menu-depth1-text" ).text( depth1 );
    $( ".menu-depth2-text" ).text( depth2 );
}
function inet_aton(dot) {
    var d = dot.split( "." );
    return ((((((+d[0])*256)+(+d[1]))*256)+(+d[2]))*256)+(+d[3]);
}

function inet_ntoa(num) {
    var d = num%256;
    for ( var i=3; i>0; i-- ) {
        num = Math.floor( num/256 );
        d = num%256 + '.' + d;
    }
    return d;
}

function getIpasTag( val ) {
    var prefix = val.substr(0, 3),
        tag = "";

    if ( prefix == "VT_" ) {
        tag += '<button class="btn blue-dark btn-xs">';

    } else if ( prefix == "ZT_" ) {
        tag += '<button class="btn blue-sharp btn-xs">';

    } else if ( prefix == "PT_" ) {
        tag += '<button class="btn green-sharp btn-xs">';

    } else {
        tag += '<i class="fa fa-question"> ';
    }
    tag += val+'</button>';

    return tag;
}



// function getBit(value, pos) {
    // return !!( value & ( 1 << pos ) );
// }

// function getXsrsToken() {
//     var _xsrf = Cookies.get( "_xsrf" ).split( "|" ),
//         xsrf = $.base64.decode(_xsrf[0]);
//     return xsrf;
// }

// 테이블 키 생성
function getTableKey( $table, ctrl ) {
    return 'tk_' + ctrl + "/" + $table.attr( "id" );
}

// 테이블 데이터 디스플레이
function showTableData($t, paging, logs) {
    var offset = (( paging.no - 1 ) % paging.blockSize ) * paging.size;
    $t.bootstrapTable( "load", logs.slice( offset, offset + paging.size ) );
    //console.log(offset + " ~ " + ( offset + paging.size ) );
}

// 테이블 컬럼 저장
function captureTableColumns( table, key ) {
    var cols = [];
    $( table ).find( "th" ).each(function( i, th ) {
        var col = $( th ).data( "field" );
        cols.push( col );
    });
    $.cookie( key, cols.join(","),  { expires: 365 } );
}

//  테이블 컬럼 복구
function restoreTableColumns( table, key ) {
    if ( $.cookie( key ) !== undefined ) {
        try {
            var h = {};
            $.map( $.cookie( key ).split( "," ), function ( col, i ) {
                h[col] = true;
                $( table ).bootstrapTable("showColumn", col);
            });

            $( table ).find( "th" ).each(function (i, th) {
                var col = $( th ).data("field");
                if ( h[col] ) {
                    $( table ).bootstrapTable("showColumn", col);
                } else {
                    $( table ).bootstrapTable("hideColumn", col);
                }
            });
        } catch(err) {
            console.log( err );
        }
    }
}

// 네비게이션 버튼 상태변경(고속 페이징)
function updateToolbarNav( $table, paging, logLength ) {
    var $toolbar = $( $table.data( "toolbar" ) ),
        offset = (( paging.no - 1 ) % paging.blockSize ) * paging.size;
    if ( logLength - offset < paging.size ) {
        $( ".btn-next", $toolbar ).prop( "disabled", true );
    } else {
        $( ".btn-next", $toolbar ).prop( "disabled", false );
    }

    if ( paging.no == 1 ) {
        $( ".btn-prev", $toolbar ).prop( "disabled", true );
    } else {
        $( ".btn-prev", $toolbar ).prop( "disabled", false );
    }
}


function beautifyMessage( obj ) {
    var r = JSON.parse( obj.message ),
        m = {
            title: "",
            body: ""
        };

    // console.log( r );
    var srcEquipType = r.equip_id.substr(0, 2).toLowerCase(),
        dstEquipType = r.targets.substr(0, 2).toLowerCase();
    if ( r.event === "speeding" ) {
        m.title = "SPEED VIOLATION !";
        if ( srcEquipType === "vt" || srcEquipType === "pt" || srcEquipType === "zt" ) {
            m.body += "<img src='/static/assets/img/obj/" + srcEquipType + ".png' height='80'>";
            m.body += r.speed + " km/h";
        }

    } else if ( r.event === "proximity" ) {
        m.title = "SHOCK !";
        if ( srcEquipType === "vt" || srcEquipType === "pt" || srcEquipType === "zt" ) {
            m.body += "<img src='/static/assets/img/obj/" + srcEquipType + ".png' height='80'>";
        }
        m.body += ' <i class="fa fa-caret-left s20"></i> <i class="fa fa-caret-right s20"></i> ';
        if ( dstEquipType === "vt" || dstEquipType === "pt" || dstEquipType === "zt" ) {
            m.body += "<img src='/static/assets/img/obj/" + dstEquipType + ".png' height='80'>";
        }
    }
    return m;
}


function updateNews() {
    $( "#pgb" ).css( "width", 11 );

    $.ajax({
        type: "GET",
        async: true,
        url:   "/news"
    }).done( function( news ) {

        // Message
        if ( news.message !== null && news.message.length > 0 && $( "#message-id-0" ).length === 0 ) {
            toastr.success( '<a href="#" id="message-id-0" class="btn-global-message" data-message-id="0">', "Mark all as read", {timeOut: 0, closeButton: false} );
        }
        $.each( news.message, function( i, r ) {
            var id = 'message-id-' + r.message_id,
                len = $( "#" + id ).length;
            if ( len > 0 ) {
                return true;
            }

            var msgHeader = '<a href="#" id="message-id-' + r.message_id + '" class="btn-global-message" data-message-id="' + r.message_id + '">',
                msgFooter = '</a>',
                timeInfo = '<div class="text-right small">' + moment(r.date).format("lll") + '</div>',
                msg = beautifyMessage( r );

            if ( r.priority === 2 ) { // Warning
                 toastr.warning( msgHeader + msg.body + msgFooter + timeInfo, msg.title );

            } else if ( r.priority === 4 ) { // Danger
                toastr.error( msgHeader + msg.body + msgFooter + timeInfo, msg.title );

            } else {
                toastr.info( msgHeader + msg.body + msgFooter + timeInfo, msg.title );
            }
        });

        // CPU
        $( "#pgb-cpu" ).css( "width", news.resource.cpu_usage );
        $( ".usage-cpu" ).text( news.resource.cpu_usage + "%" );

        // Memory
        var memUsage = news.resource.mem_used / news.resource.mem_total * 100,
            memTotal = news.resource.mem_total / 1024 / 1024 / 1024,
            memUsed = news.resource.mem_used / 1024 / 1024 / 1024;

        $( "#pgb-mem" ).css( "width", memUsage.toFixed(1) );
        $( ".usage-mem" ).text( memUsage.toFixed(1) + "%" );
        $( ".used-mem" ).text( memUsed.toFixed(1) + " GB" );
        $( ".total-mem" ).text( memTotal.toFixed(1) + " GB" );


        // Disk
        var diskUsage = news.resource.disk_used / news.resource.disk_total * 100,
            diskTotal = news.resource.disk_total / 1024 / 1024 / 1024,
            diskUsed = news.resource.disk_used / 1024 / 1024 / 1024;
        $( "#pgb-disk" ).css( "width", diskUsage );
        $( ".usage-disk" ).text( diskUsage.toFixed(1) + "%" );
        $( ".used-disk" ).text( diskUsed.toFixed(1) + " GB" );
        $( ".total-disk" ).text( diskTotal.toFixed(1) + " GB" );

        // Clock
        $( ".system-clock" ).text( moment( news.time ).format( "lll (Z)" ) );
    });
}