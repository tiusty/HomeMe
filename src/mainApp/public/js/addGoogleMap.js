function initMap() {
	var mapDiv = document.getElementById('map');
	var mpa = new google.maps.Map(mapDiv, {
		center: {lat:42.3601, lng: -71.0589},
		zoom: 8
	});
}
