function rankFormatter( val, row, idx ) {
    var btn_class = 'btn-default';
    if (val == 1) {
        btn_class = 'blue';
    }

    return '<button class="btn ' + btn_class + ' btn-xs">' + val + '</button>';
}

function orgGroupNameFormatter( val, row, idx ) {
    var groupName = row.group_name;
    if ( row.item.endsWith("/0") ) {
        groupName = '<span class="font-grey-silver">Default</span>';
    }
    return row.org_name + '<i class="fa fa-angle-right mlr10"></i>' + groupName;
}