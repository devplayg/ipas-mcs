$(function() {

//$("#modal-sensor-add").modal("show");
//$("#form-sensor-add select[name=Type1]").val(1).selectpicker( "refresh" );
//$("#form-sensor-add input[name=Name]").val("No name");
//$("#form-sensor-add input[name=Hostname]").val("10.10.10.10");
//$("#form-sensor-add input[name=Port]").val(4000);

    /**
     * 1. 초기화
     *
     */

    // Tree
    var $tree = $( "#tree-assets" );
    $tree.jstree({
        "core" : {
            "data" : {
                "url" : "/assetclass/1/descendants/0",
            },
            "check_callback" : function ( op, node, node_parent, node_position, more ) {
                var ids = $tree.jstree( true ).get_selected();
                console.log(node);
                if ( op === "create_node" ) {
                    $( "#modal-sensor-add" ).modal( "show" );

                } else if ( op === "rename_node" ) {  //  can be 'create_node', 'rename_node', 'delete_node', 'move_node', 'copy_node' or 'edit'
                    $.ajax({
                        type: "patch",
                        async: true,
                        url: "/assets/" + node.original.AssetId,
                        data: {
                            Name: node_position
                        }
                    }).done( function( rows ) {
                    }).always( function() {
//                    refreshTree();
                    });
                } else if ( op === "delete_node" ) {
                    delete_node( node );
                }

                return ( op === "rename_node" ) ? true : false;
            },
            "multiple" : false,
            "animation" : 1,
//        'themes': {
//            'name': 'proton',
            'responsive': true
//        }
        },
//         "contextmenu": {
//             "items": function( $node ) {
//                 var type = this.get_type( $node );
//
//                 if ( type == "type_0" ) { // Root
//                     return {
//                         "Create": {
//                             "separator_before": true,
//                             "separator_after": true,
//                             "label": "Create sensor",
//                             "action": function ( obj ) {
//                                 $node = $tree.jstree( true ).create_node($node);
// //                            $tree.jstree( true ).edit( $node );
// //                            $node = $tree.create_node( $node );
// //                            $tree.edit( $node );
//                             }
//                         }
//                     };
//
//                 } else if ( type == "type_1" ) { // Companies
//                     return {
//                         "Create": {
//                             "separator_before": false,
//                             "separator_after": true,
//                             "label": "Create groups",
//                             "icon": "fa fa-plus",
//                             "action": function( obj ) {
// //                            $node = tree.create_node( $node );
// //                            tree.edit( $node );
//                             }
//                         },
//                         "Rename": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Rename",
//                             "icon": "fa fa-tag",
//                             "action": function( obj ) {
//                                 $tree.jstree( true ).edit( $node );
//                             }
//                         },
//                         "Edit": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Edit",
//                             "icon": "fa fa-edit",
//                             "action": function( obj ) {
// //                            $tree.jstree( true ).edit( $node );
//                             }
//                         },
//                         "Remove": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Remove",
//                             "icon": "fa fa-trash-o",
//                             "action": function( obj ) {
//                                 $tree.jstree( true ).delete_node( $node );
//                             }
//                         }
//                     };
//                 } else if (type == "type_2" ) { // Departments
//                     return {
//                         "Create": {
//                             "separator_before": false,
//                             "separator_after": true,
//                             "label": "Create network",
//                             "icon": "fa fa-plus",
//                             "action": function( obj ) {
// //                            $node = tree.create_node( $node );
// //                            tree.edit( $node );
//                             }
//                         },
//                         "Rename": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Rename",
//                             "icon": "fa fa-trash-o",
//                             "action": function( obj ) {
//                                 $tree.jstree( true ).edit( $node );
// //                            tree.edit( $node );
//                             }
//                         },
//                         "Edit": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Edit",
//                             "icon": "fa fa-edit",
//                             "action": function( obj ) {
// //                            alert(3);
//                             }
//                         },
//                         "Remove": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Remove",
//                             "icon": "fa fa-trash-o",
//                             "action": function( obj ) {
//                                 $tree.jstree( true ).delete_node( $node );
//                             }
//                         }
//                     };
//
//                 } else if (type == "type_4" ) { // Network
//                     return {
//                         "Rename": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Rename",
//                             "icon": "fa fa-edit",
//                             "action": function( obj ) {
//                                 $tree.jstree( true ).edit( $node );
//                             }
//                         },
//                         "Edit": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Edit",
//                             "icon": "fa fa-edit",
//                             "action": function( obj ) {
// //                            alert(3);
//                             }
//                         },
//                         "Remove": {
//                             "separator_before": false,
//                             "separator_after": false,
//                             "label": "Remove",
//                             "icon": "fa fa-trash-o",
//                             "action": function( obj ) {
//                                 $tree.jstree( true ).delete_node( $node );
//                             }
//                         }
//                     };
//                 }
//             }
//         },
        "types" : {
            "default": {
                icon: "fa fa-folder icon-state-warning"
            },
            "type_0": {
                //icon: "fa fa-ban"
            },
            "type_1": {
                icon: "fa fa-building-o"
            },
            "type_2": {
                icon: "fa fa-folder font-blue-sharp"
            },
            "type_4": {
                icon: "fa fa-sitemap"
            },
        },

        "plugins" : [
            "types", "state", "sort", "contextmenu"
        ]
    }).on( "loaded.jstree", function() {
        // $(this).jstree( "open_all");

    }).on( "changed.jstree", function( e, obj ) {
        // console.log("changed.jstree: " + obj.action);
        // obj.
        console.log(obj.node);
        $( "#table-assets" ).bootstrapTable( "load", [] );
        if ( obj.action == "select_node" ) {
            var assetId = obj.node.original.AssetId;
            if (assetId == "j1_1" ) {
                assetId = 0;
            }

            $.ajax({
                type: "get",
                async: true,
                url: "/assets/" + assetId + "/children"
            }).done( function( rows ) {
                if ( rows !== null ) {
                    $( "#table-assets" ).bootstrapTable( "load", rows );
                }
            });
        }
    });


//$.ajax({
//    type: "GET",
//    async: true,
//    url: "/assetclass/1/descendants/0"
//}).done(function(org) {
//    $tree.jstree(true).settings.core.data = org;
//    $tree.jstree(true).refresh(-1);

//}).fail(function() {

//}).always(function() {
//});


//$.ajax({
//    type  : "get",
//    async : true,
//    url   : "/assets/class/1/tree"
//}).done( function( result ) {
//    $( "#tree-assets" ).jstree( true ).settings.core.data = result;
//    $( "#tree-assets" ).jstree( true ).refresh(-1);
//    $( "#tree-assets" ).jstree( "open_all", -1);
//}).fail(function() {
//}).always(function() {
//     $( "#tree-assets" ).jstree( "open_all", -1);
//});


    /**
     * 2. 이벤트
     *
     */

    $( ".btn-tree-refresh" ).click( function( e ) {
        e.preventDefault();
        refreshTree();
    });
    $( ".btn-tree-expand" ).click( function( e ) {
        e.preventDefault();
        $tree.jstree( "open_all" );
    });
    $( ".btn-tree-collapse" ).click( function( e ) {
        e.preventDefault();
        $tree.jstree( "close_all" );
    });

//$( "#tree-assets').on( "changed.jstree", function (e, data) {
//    console.log(3);
//    $( ".btn-page').text(1);
//    $( "#form-filter input[name=offset]').val(0);

//    //if (defined(data.node.original)) {
//    if (data.node != undefined) {
//        // Nav. text
//        var nav = data.node.original.text_origin;
//        if (data.node.parents.length == 2) {
//            nav = data.instance.get_node(data.node.parent).original.text_origin + " <span class='mlr10'> &gt; </span> " + nav;
//        }
//        if (data.node.parents.length == 3) {
//            nav = data.instance.get_node(data.node.parent).original.text_origin + " <span class='mlr10'> &gt; </span> " + nav;
//            nav = data.instance.get_node(data.instance.get_node(data.node.parent).parent).original.text_origin + " <span class='mlr10'> &gt; </span> " + nav;
//        }
//        $( ".btn-org-tag').html(nav);
//        var param = {
//            depth: data.node.parents.length,
//            id: data.node.id,
//            keyword: $.trim($( "#form-filter input[name=keyword]').val())
//        };

//        var url = '?mod=agent&act=procGetAgent&with_version=1&' + filter_query + '&' + $.param(param);
//        $( "#table_agent').bootstrapTable( "refresh', {url: url} );
//    } else {
//        started = 1;
//    }
//});






    /**
     * 3. 함수
     *
     */

    function refreshTree() {
        $tree.jstree( true ).refresh( -1 );
    }


    var delete_node = function( node ) {
        swal({
            title: "Are you sure?",
            text: "You will not be able to recover this imaginary file!",
            type: "warning",
            showLoaderOnConfirm: true,
            showCancelButton: true,
            confirmButtonText: "Yes, delete it!",
            cancelButtonText: "No, keep it",
            confirmButtonColor: "#E7505A",
        }).then(function() {
            $.ajax({
                type: "DELETE",
                url: "/assets/" + node.original.AssetId,
                async: true,
                beforeSend: function(xhr) {
                    xhr.setRequestHeader("X-CSRFToken", getXsrsToken());
                },
            }).done( function( result ) {
                if ( result.State ) {
//                $tree.jstree( true ).refresh( -1 );
                } else {
                    swal( "Fail", result.Message, "error" );
                }
            }).fail( function() {
            }).always( function() {
                refreshTree();
            });
        })
    }


});

