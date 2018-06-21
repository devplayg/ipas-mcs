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
    var $tree = $( "#tree-assets" ),
        Root = 0,
        Org = 1,
        Group = 2,
        $ipasTable = $( "#table-ipas" );

    $tree.jstree({
        "plugins" : [
            "types", "state", "sort"
        ],

        "core" : {
            "data" : {
                "url" : "/assetclass/1/root/0",
            },
            "multiple": false
        },
        "types" : {
            1: {
                icon: "fa fa-building-o"
            },
            2: {
                icon: "fa fa-folder font-blue-sharp"
            }
        },

    }).on( "changed.jstree", function( e, obj ) {
        // console.log(selected);
    }).on( "select_node.jstree", function( e, obj ) {
        updateTreeButtons();
        if ( obj.node.original.type > 0 ) {
            getLeafNodes( obj.node );
        }
        showIpasList( obj.node );
    });


    $ipasTable.on('check.bs.table', function (e, row) {
        updateTableButtons( $( this ) );

    }).on('uncheck.bs.table', function (e, row) {
        updateTableButtons( $( this ) );

    }).on('check-all.bs.table', function (e) {
        updateTableButtons( $( this ) );

    }).on('uncheck-all.bs.table', function (e) {
        updateTableButtons( $( this ) );
    }).on( "refresh.bs.table", function() {
        updateTableButtons( $ipasTable );

    });


    function updateTableButtons( $table ) {
        var selectedNode = $tree.jstree( true ).get_selected();
            node = $tree.jstree( true ).get_node( selectedNode[0] );

        if ( ! node ) {
            return;
        }

        var selectedIpas = $table.bootstrapTable( "getSelections" );
        if ( node.original.type > 0 && selectedIpas.length > 0 ) {
            $( ".btn-asset-classify" ).removeClass( "hide" );
        } else {
            $( ".btn-asset-classify" ).addClass( "hide" );
        }
    }

    function getLeafNodes( node ) {
        var list,
            parentName = '';

        if ( node.original.type == Org ) {
            list = node.children;
            parentName = node.text;

        } else if ( node.original.type == Group ) {
            var parent = $tree.jstree( true ).get_node( node.parent );
            list = parent.children;
            parentName = parent.text;
        } else {
            return;
        }
        $( ".ast-group" ).empty();

        for (var i=0; i<list.length; i++) {
            var n = $tree.jstree( true ).get_node( list[i] );

            $( ".ast-group" ).append('<li><a href="#" class="btn-asset-tag" data-id="' + n.id + '">' + parentName + ' > ' + n.text + '</a></li>');;
            // console.log(n.id + " / " + n.text);
        }
    }

    $( document ).on( "click",".btn-asset-tag", function(e) {
        console.log(3);
        var group_id = $( this ).data( "id" );
            // selected = $ipasTable.bootstrapTable( "getSelections" );

        var list = $ipasTable.bootstrapTable( "getSelections" ).map(function(r) {
            return r.org_id + "/" + r.equip_id;
        });
        console.log(list);

        $.ajax({
            type:  "PATCH",
            async: true,
            url:   "/ipasgroup/" + group_id,
            data: {
                list: list
            }
        }).done( function( result ) {
            if ( result.state ) {
                $ipasTable.bootstrapTable( "removeAll" );
                // updateTableButtons( $ipasTable );
                $ipasTable.bootstrapTable( "refresh" );
            }
        });

    });



    /**
     * 2. 이벤트
     *
     */
    $( ".btn-tree-refresh" ).click( function( e ) {
        $tree.jstree( true ).refresh( -1 );
    });

    $( ".btn-tree-expand" ).click( function( e ) {
        var root = $tree.jstree( true ).get_node("j1_1").state;
        if ( root.opened ) {
            $tree.jstree( "close_all" );
        } else {
            $tree.jstree( "open_all" );
        }
    });

    $( ".btn-tree-collapse" ).click( function( e ) {
        $tree.jstree( "close_all" );
    });

    $( "#modal-asset-add" )
        .on( "shown.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $( "input[name=name]", $form ).focus();

        }).on( "hidden.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $form.get( 0 ).reset();
            $( ".form-option" ).addClass( "hide" );
        });

    $( "#modal-asset-edit" )
        .on( "shown.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $( "input[name=name]", $form ).focus().select();

        }).on( "hidden.bs.modal", function () {
            var $form = $( this ).find( "form" );
            $form.get( 0 ).reset();
            $( ".form-option" ).addClass( "hide" );
        });


    $( ".btn-asset-add" ).click(function() {
        var selected = $tree.jstree( true ).get_selected(),
            node = $tree.jstree( true ).get_node( selected[0] );

        var target ;
        if ( node.original.type === Root ) {
            target = '<i class="fa fa-building-o"></i> ' + felang[ "org" ];
            $( ".form-option-org" ).removeClass( "hide" );
        } else if ( node.original.type === Org ) {
            target = '<i class="fa fa-folder-o"></i> ' + felang[ "group" ];
            $( ".form-option-group" ).removeClass( "hide" );
        } else {
            return;
        }

        $( "#form-asset-add input[name=parent_id]" ).val( node.original.asset_id );
        $( "#form-asset-add input[name=type1]" ).val( node.original.type + 1 );
        $( "#form-asset-add .target" ).html( target );
        $( "#form-asset-add .name" ).html(
            node.original.name.length > 0 ? '<i class="fa fa-chevron-left"1></i> <label class="label label-info">' + node.original.name + '</label>'
                                      : ""
        );

        $( "#modal-asset-add" ).modal( "show" );
    });


    $( "#form-asset-add" ).submit(function( e ) {
        e.preventDefault();
        var $form = $( this );

        $.ajax({
            type: "POST",
            async: true,
            url: "/assets",
            data: $form.serialize()
        }).done( function( result ) {
            if ( result.state ) {
                $( "#modal-asset-add" ).modal( "hide" );
                $tree.jstree( true ).refresh( -1 );
            } else {
                swal(result.message, "", "error");
            }
        });
    });


    $( ".btn-asset-edit" ).click(function() {
        var selected = $tree.jstree( true ).get_selected(),
            node = $tree.jstree( true ).get_node( selected[0] );

        if ( node.original.type === Root ) {
            return;
        }

        $.ajax({
            type: "GET",
            async: true,
            url: "/assets/" + node.original.asset_id,
        }).done( function( result ) {
            if ( result.state ) {
                var target ;
                if ( node.original.type === Org ) {
                    target = '<i class="fa fa-building-o"></i> ' + felang[ "org" ];
                    $( ".form-option-org" ).removeClass( "hide" );
                } else if ( node.original.type === Group ) {
                    target = '<i class="fa fa-folder-o"></i> ' + felang[ "group" ];
                    $( ".form-option-group" ).removeClass( "hide" );
                } else {
                    return;
                }

                $( "#form-asset-edit .target" ).html( target );
                $( "#form-asset-edit input[name=name]" ).val( result.data.name );
                $( "#form-asset-edit input[name=code]" ).val( result.data.code );
                $( "#form-asset-edit input[name=asset_id]" ).val( node.original.asset_id );
                $( "#modal-asset-edit" ).modal( "show" );

            } else {
                swal(result.message, "", "error");
            }
        });
    });


    $( ".btn-asset-remove" ).click(function() {
        swal({
            title: felang[ "msg.confirm_delete" ],
            type: "warning",
            showCancelButton: true,
            confirmButtonClass: "btn red",
            confirmButtonText: felang[ "yes" ],
            cancelButtonText: felang[ "no" ]
        }).then((result) => {
            if (result.value) {
                var list = $tree.jstree( true ).get_selected().map( Number );
                $.ajax({
                    type: "POST",
                    async: true,
                    url: "/assets/delete",
                    data: {
                        asset_id_list: list
                    }
                }).done( function( result ) {
                    if ( result.state ) {
                        $tree.jstree( true ).refresh( -1 );
                    } else {
                        swal(result.message, "", "error");
                    }
                }).always(function() {
                    $ipasTable.bootstrapTable( "removeAll" );
                    // updateTableButtons( $ipasTable );
                    $ipasTable.bootstrapTable( "refresh" );
                });
            }
        });
    });


    $( "#form-asset-edit" ).submit(function( e ) {
        e.preventDefault();

        var $form = $( this );

        $.ajax({
            type: "patch",
            async: true,
            url: "/assets/" + $( "input[name=asset_id]", $form ).val(),
            data: $form.serialize()
        }).done( function( result ) {
            if ( result.state ) {
                $( "#modal-asset-edit" ).modal( "hide" );
                $tree.jstree( true ).refresh( -1 );
            } else {
                swal(result.message, "", "error");
            }
        });
    });



    /**
     * 3. 함수
     *
     */

    function updateTreeButtons() {
        var selected = $tree.jstree( true ).get_selected();
        if ( selected.length < 1 ) {
            $( ".btn-asset-manage" ).addClass( "hide" );
        } else if ( selected.length == 1 ) {
            $( ".btn-asset-manage" ).addClass( "hide" );
            var node =  $tree.jstree( true ).get_node( selected[0] );
            if ( node.original.type == Root ) {
                $( ".btn-asset-remove" ).addClass( "hide" );
            } else {
                $( ".btn-asset-remove" ).removeClass( "hide" );
            }
            if ( node.original.type == Root || node.original.type == Org ) {
                $( ".btn-asset-add" ).removeClass( "hide" );
            }

            if ( node.original.type == Org || node.original.type == Group ) {
                $( ".btn-asset-edit" ).removeClass( "hide" );
            }
        } else {
            $( ".btn-asset-manage" ).addClass( "hide" );
        }
    }


    function showIpasList( selected ) {
        var url;

        if ( selected.original.type == Root ) {
            url = "/ipasorg/0";

        } else if ( selected.original.type == Org ) {
            url = "/ipasorg/" + selected.id;

        } else if ( selected.original.type == Group ) {
            url = "/ipasgroup/" + selected.id;
        }

        $ipasTable.bootstrapTable( "removeAll" );
        // updateTableButtons( $ipasTable );
        // $ipasTable.bootstrapTable('selectPage', 1);
        $ipasTable.bootstrapTable( "refresh", { url: url,  pageNumber: 1 });
    }


});