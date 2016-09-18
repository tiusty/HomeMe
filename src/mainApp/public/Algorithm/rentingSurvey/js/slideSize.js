$( function() {
    $( "#size" ).slider({
        range: true,
        min: 0,
        max: 500,
        values: [ 75, 300 ],
        slide: function( event, ui ) {
            $( "#amountSize" ).val( ui.values[ 0 ] + " square feet to " + ui.values[ 1 ] + " square feet");
            $( "#amountSizeLow" ).val( ui.values[ 0 ]);
            $( "#amountSizeHigh" ).val(ui.values[ 1 ]);
        }
    });
    $( "#amountSize" ).val($( "#size" ).slider( "values", 0 ) +
        " square feet to " + $( "#size" ).slider( "values", 1 ) + " square feet" );
    $( "#amountSizeLow" ).val( $("#size").slider("values", 0));
    $( "#amountSizeHigh" ).val( $("#size").slider("values", 1));
} );