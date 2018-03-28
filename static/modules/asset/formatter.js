function assetsHostnameFormatter(val, row, idx) {
    if ( row.Type1 == 1 ) {
        return val + " : " + row.Port;
    } else if ( row.Type1 == 2 ) {
        return null;

    } else if ( row.Type1 == 4 ) {
        return val + " / " + row.Cidr;
    }

    return val;
}


function assetsNameFormatter( val, row, idx ) {
    var icon = "";
//        name = "";
    if ( row.Type1 == 1 ) {
        icon = '<i class="fa fa-server"></i>';
//        name = 'Sensor';

    } else if ( row.Type1 == 2 ) {
        icon = '<i class="fa fa-clone"></i>';
//        name = 'Group';

    } else if ( row.Type1 == 4 ) {
        icon = '<i class="fa fa-sitemap"></i>';
//        name = 'Network';
    }

    return icon + " " + val;
}

