$( function() {
    $( "#size" ).slider({
        range: true,
        min: 0,
        max: 500,
        values: [ 75, 300 ],
        slide: function( event, ui ) {
            $( "#amountSize" ).val( ui.values[ 0 ] + " square feet to " + ui.values[ 1 ] + " square feet");
        }
    });
    $( "#amountSize" ).val($( "#size" ).slider( "values", 0 ) +
        " square feet to " + $( "#size" ).slider( "values", 1 ) + " square feet" );
} );