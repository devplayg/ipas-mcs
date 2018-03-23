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
    return m.format("YYYY-MM-DD HH:mm:ss");
}

function numberFormatter( val, row, idx ) {
    return val.toLocaleString();
}

function equipIdFormatter( val, row, idx ) {
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

function targetEquipIdFormatter( val, row, idx ) {
    var list = val.split(","),
        tags = '';

    for ( var i=0; i<list.length; i++ ) {
        tags += equipIdFormatter(list[i]);
    }

    return tags;

}