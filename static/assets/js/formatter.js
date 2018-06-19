function ipWithTagFormatter( val, row, idx ) {
    var ip = inet_ntoa( val );
    if ( row.SrcIpCard.length > 0 ) {
        var ipcard = $.parseJSON( row.SrcIpCard );
        var head = '<button class="btn default btn-xs pull-right ipcard" data-toggle="modal" data-target="#modal-ipcard" data-ip="' + ip + '"><i class="fa fa-id-card-o"></i> ',
            body = ipcard.Name,
            foot = '</button>';
        return head + body + foot + ip;
    } else {
        var head = '<a href="#"><i class="fa fa-tag pull-right ipcard" data-toggle="modal" data-target="#modal-ipcard" data-ip="' + ip + '"></i> ',
            body = '',
            foot = '</a>';
        return head + body + foot + ip;
    }
}

function ip2intFormatter( val, row, idx ) {
    return inet_aton(val);
}

function int2ipFormatter( val, row, idx ) {
    return inet_ntoa(val);
}

function dateFormatter( val, row, idx ) {
    var m = moment( val );
    // return '<span class="">' + m.format( "MMMM D YYYY, hh:mm:ss" ) + '</span>';
    return '<span class="">' + m.format() + '</span>';
}

function shortDateFormatter( val, row, idx ) {
    var m = moment( val ),
        prefix = '<span class="tooltips" data-container="body" data-placement="top" data-original-title="' + m.format( "MMMM D YYYY, h:mm:ss a" ) + '" title="' + m.format( "MMMM D YYYY, h:mm:ss a" ) + '">',
        suffix = '</span>';
    return prefix + m.format("HH:mm") + suffix;
}

function numberFormatter( val, row, idx ) {
    return val.toLocaleString();
}


function ipasEquipIdFormatter( val, row, idx ) {
    var header = '',
        body = '',
        footer = '';

    header += '<a href="#" data-toggle="modal" data-target="#modal-ipas-report" data-org-id="' + row.org_id + '" data-equip-id="' + row.equip_id + '" data-encoded="' + encodeURI( JSON.stringify( row ) ) + '" >';
    body += getIpasTag( row.equip_id );
    return header + body + footer;
}

function orgNameFormatter( val, row, idx ) {
    // return '<span class="tooltips" data-container="body" data-placement="top" data-original-title="' + row.org_id + '">' + row.org_name + '</span>'
    return '<span class="tooltips" title="Org ID: ' + row.org_id + '">' + row.org_name + '</span>'
}

function groupNameFormatter( val, row, idx ) {
    if ( row.group_id == 0 ) {
        return;
    }
    return row.group_name;
}

function orgGroupNameFormatter( val, row, idx ) {
    var orgName = orgNameFormatter( row.org_id, row, null );
    if ( row.group_id == 0 ) {
        return;
    }
    return orgName + " / " + row.group_name;
}



function ipasEventRowStyle( row, idx ) {
    if ( row.event_type == ProximityEvent ) {
        return {
            classes: "row-danger"
        };
    }
    // if ( row.shock_count >= 10 ) {
    //     return {
    //         classes: "row-danger"
    //     };
    // } else if ( row.shock_count >= 8 ) {
    //     return {
    //         classes: "row-warning"
    //     };
    // }

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