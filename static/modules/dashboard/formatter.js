function rankFormatter( val, row, idx ) {
    var btn_class = 'btn-default';
    if (val == 1) {
        btn_class = 'blue';
    }

    return '<button class="btn ' + btn_class + ' btn-xs">' + val + '</button>';
}

function dashboardOrgGroupNameFormatter( val, row, idx ) {
    var groupName = row.group_name;
    if ( row.item.endsWith("/0") ) {
        groupName = '<span class="font-grey-silver">Default</span>';
    }

    var asset = row.item.split( "/", 2 );
    var param = {
        orgId: asset[0],
        groupId: asset[1]
    };
    var prefix = '<a href="#" style="color: inherit; " data-toggle="modal" data-target="#modal-ipaslog" data-query="'+ $.param(param) + '">',
        suffix = '</a>';
    return prefix + row.org_name + '<i class="fa fa-angle-right mlr10"></i>' + groupName + suffix;
}