function shockCountFormatter(val, row, idx) {
    if ( row.shock_count >= 10 ) {
        return '<button class="btn red btn-xs">' +  val + '</button>';

    } else if ( row.shock_count >= 8 ) {
        return '<button class="btn btn-warning btn-xs">' +  val + '</button>';
    }
}


function ipaslogEquipIdFormatter(val, row, idx) {
    var header = '',
        body = '',
        footer = '';

    header += '<a href="#" data-toggle="modal" data-target="#modal-ipas-report" data-equip-id="' + row.equip_id + '" data-encoded="' + encodeURI(JSON.stringify(row)) + '" >';
    body = getIpasTag( row.equip_id );
    return header + body + footer;
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
        color = ['red-mint', 'red-mint', 'yellow-crusta', 'yellow-crusta', 'blue', 'blue'],
        level = Math.ceil(val / 2);

    if ( level > 6 ) {
        level = 6;
    }

    return '<span class="font-' + color[level-1] + '">' + c.repeat( level ) + '</span>';
}


function ipaslogLocationFormatter(val, row, idx) {
    var loc = '';
    loc += '<a href="#" class="tooltips"  data-container="body" data-placement="top" data-original-title="Tooltip in top" data-toggle="modal" data-target="#modal-ipas-map" data-latitude="' + row.latitude + '" data-longitude="' + row.longitude + '"><i class="fa fa-map-marker fa-1x"></i></a>';
    // loc += '<small class="ml5">' + row.latitude + ", " + row.longitude + '</small>';

    return loc;
}


function ipaslogEventTypeFormatter(val, row, idx) {
    if (val === 1) {
        return felang[ "ipas.start" ] + ' <i class="fa fa-power-off pull-right"></i>';

    } else if (val === 2) {
        return felang[ "shock" ] + ' <span class="pull-right"><i class="fa fa-bolt"></i></span>';

    } else if (val === 3) {
        return felang[ "speeding" ] + ' <i class="fa fa-long-arrow-up pull-right"></i>';

    }
}