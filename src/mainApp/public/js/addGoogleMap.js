function initMap() {
	var mapDiv = document.getElementById('map');
	var map = new google.maps.Map(mapDiv, {
		center: {lat:42.3601, lng: -71.0589},
		zoom: 8
	});
	distanceMatrix();
}
