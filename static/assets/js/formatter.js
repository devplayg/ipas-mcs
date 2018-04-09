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

function numberFormatter( val, row, idx ) {
    return val.toLocaleString();
}

