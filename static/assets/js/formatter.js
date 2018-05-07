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
    // return m.format("YYYY-MM-DD HH:mm:ss");
    return '<span class="">' + m.format("YYYY-MM-DD HH:mm:ss") + '</span>';
}

function shortDateFormatter( val, row, idx ) {
    var m = moment( val ),
        prefix = '<span class="tooltips" data-container="body" data-placement="top" data-original-title="' + m.format("YYYY-MM-DD HH:mm:ss") + '">',
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

    header += '<a href="#" data-toggle="modal" data-target="#modal-ipas-report" data-equip-id="' + row.equip_id + '" data-encoded="' + encodeURI(JSON.stringify(row)) + '" >';
    body += getIpasTag( row.equip_id );
    return header + body + footer;
}

function groupNameFormatter( val, row, idx ) {
    if ( row.group_id == 0 ) {
        return;
    }
    return val;
}