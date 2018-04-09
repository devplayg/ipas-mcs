/**
 *
 */

// Ajax
var IpasStatusStart = 1,
    IpasStatusShock = 2,
    IpasStatusSpeeding = 3,
    IpasStatusProximity = 4;

$.ajaxSetup({cache:false});
var ajax = $.ajax;
$.extend({
    ajax: function(url, options) {
        if (typeof url === 'object') {
            options = url;
            url = undefined;
        }
        options = options || {};
        url = options.url;
        var xsrftoken = $('meta[name=_xsrf]').attr('content');
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

// Mask
$( ".mask-yyyymmddhhii" ).mask( "0000-00-00 00:00" );
$( ".mask-ipv4-cidr" ).mask( "099.099.099.099/09" );
$( ".mask-port" ).mask( "09999" );
$( ".mask-0999" ).mask( "0999" );
$( ".mask-099" ).mask( "099" );

// Language
$( ".lang-changed" ).click( function() {
    var $e = $( this );
    var lang = $e.data( "lang" );
    $.cookie( "lang", lang, { path: "/", expires: 365} );
    window.location.reload();
});


/**
 * Functions
 */
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
function showTableData($t, logs, paging) {
    var offset = (( paging.no - 1 ) % paging.blockSize ) * paging.size;
    $t.bootstrapTable( "load", logs.slice( offset, offset + paging.size ) );
    console.log(offset + " ~ " + ( offset + paging.size ) );
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
        var h = {};
        $.map(  $.cookie( key ).split( "," ), function( col, i ) {
            h[ col ] = true;
            $( table ).bootstrapTable( "showColumn", col );
        });

        $( table ).find( "th" ).each(function( i, th ) {
            var col = $( th ).data( "field" );
            if ( h[ col ] ) {
                $( table ).bootstrapTable( "showColumn", col );
            } else {
                $( table ).bootstrapTable( "hideColumn", col );
            }
        });
    }
}


