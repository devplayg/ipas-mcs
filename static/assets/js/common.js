/**
 *
 */

// Event types
var StartupEvent = 1,
    ShockEvent = 2,
    SpeedingEvent = 3,
    ProximityEvent = 4;

// Equipment types
var PedestrianTag = 1, // Pedestrian tag
    ZoneTag = 2, // Zone tag
    VehicleTag = 4; // Vehicle tag

var StartupColor = "#e1e5ec", // blue-chambray
    ShockColor = "#3598DC", // blue
    SpeedingColor = "#f4902f", // warning(yellow) // #f4902f
    ProximityColor = "#E7505A"; // red

// Notification
toastr.options = {
    "closeButton": true,
    "debug": false,
    "newestOnTop": true,
    "progressBar": false,
    "positionClass": "toast-top-right",
    "preventDuplicates": true,
    "onclick": null,
    "showDuration": "300",
    "hideDuration": "1000",
    "timeOut": "10000",
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
}, 60 * 1000);



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

    $.ajax({
        type: "GET",
        async: true,
        url:   "/message/gotit/" + messageId
    }).done( function() {

    });
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

            $(table).find("th").each(function (i, th) {
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

function updateNews() {
    $( "#pgb" ).css( "width", 11 );
    $.ajax({
        type: "GET",
        async: true,
        url:   "/news"
    }).done( function( news ) {

        // Message

        $.each( news.message, function( i, r ) {
            // console.log(r);
            var msgHeader = '<a href="#" class="btn-global-message" data-message-id="' + r.message_id + '">',
                msgFooter = '</a>',
                timeInfo = '<div class="text-right small">' + moment(r.date).format("lll") + '</div>';
            if ( r.priority === 1 ) { // Info
                toastr.info( msgHeader + r.message + msgFooter + timeInfo, "INFORMATION" );
            } else if ( r.priority === 3 ) { // Warning
                 toastr.warning( msgHeader + r.message + msgFooter + timeInfo, "WARNING" );
            } else if ( r.priority === 5 ) { // Danger
                toastr.error( msgHeader + r.message + msgFooter + timeInfo, "ERROR" );
            }
        });

        // CPU
        $( "#pgb-cpu" ).css( "width", news.resource.cpu_usage );
        $( ".usage-cpu" ).text( news.resource.cpu_usage + "%" );

        // Memory
        var memUsage = news.resource.mem_used / news.resource.mem_total * 100;
        $( "#pgb-mem" ).css( "width", memUsage.toFixed(1) );
        $( ".usage-mem" ).text( memUsage.toFixed(1) + "%" );

        // Disk
        var diskUsage = news.resource.disk_used / news.resource.disk_total * 100;
        $( "#pgb-disk" ).css( "width", diskUsage );
        $( ".usage-disk" ).text( diskUsage.toFixed(1) + "%" );

        // Clock
        $( ".system-clock" ).text( moment( news.time ).format( "LLL" ) );
    });
}