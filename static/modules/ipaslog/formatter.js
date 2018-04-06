function shockCountFormatter(val, row, idx) {
    if ( row.shock_count >= 10 ) {
        return '<button class="btn red btn-xs">' +  val + '</button>';

    } else if ( row.shock_count >= 8 ) {
        return '<button class="btn btn-warning btn-xs">' +  val + '</button>';
    }
}


function rowStyle(row, idx) {
    if ( row.shock_count >= 10 ) {
        return {
            classes: "row-danger"
        };
    } else if ( row.shock_count >= 8 ) {
        return {
            classes: "row-warning"
        };
    }

    return {};
}

function snrFormatter(val, row, idx) {
    var c  = '\u2759',
        color = ['red', 'red', 'yellow', 'yellow', 'green', 'green'],
        level = Math.ceil(val / 2);

    if ( level > 6 ) {
        level = 6;
    }

    return '<span class="font-' + color[level-1] + '">' + c.repeat( level ) + '</span>';
}


function ipaslogLocationFormatter(val, row, idx) {
    return '<a href="#" data-toggle="modal" data-target="#modal-ipas-map" data-latitude="' + row.latitude + '" data-longitude="' + row.longitude + '"><i class="fa fa-map-o"></i></a>';
}