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
        "plugins" : [
            "types", "state", "sort"
        ],

        "core" : {
            "data" : {
                "url" : "/assetclass/1/root/0",
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
            // "multiple" : false,
            // "animation" : 1,
//        'themes': {
//            'name': 'proton',
//             'responsive': true
//        }
        },
        "types" : {
            "default": {
                icon: "fa fa-folder icon-state-warning"
            },
            "type_0": {
                //icon: "fa fa-ban"
            },
            1: {
                icon: "fa fa-building-o"
            },
            2: {
                icon: "fa fa-folder font-blue-sharp"
            },
            "type_4": {
                icon: "fa fa-sitemap"
            },
        },

    }).on( "loaded.jstree", function() {
        // $(this).jstree( "open_all");

    }).on( "changed.jstree", function( e, obj ) {
        if ( obj.action == "select_node" ) {
            console.log(obj.node);
        }
    });


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

