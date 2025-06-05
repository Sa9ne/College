import 'dart:convert';
import 'package:frontend/models/car.dart';
import 'package:frontend/models/order.dart';
import 'package:http/http.dart' as http;

Future<List<Car>> fetchCars() async {
  final response = await http.get(Uri.parse('http://localhost:8080/Catalog'));

  if (response.statusCode == 200) {
    final List<dynamic> jsonData = json.decode(response.body);
    return jsonData.map((carJson) => Car.fromJson(carJson)).toList();
  } else {
    throw Exception('Failed to load cars');
  }
}

Future<int?> addClient(Map<String, dynamic> clientData) async {
  final response = await http.post(
    Uri.parse('http://localhost:8081/AddClient'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode(clientData),
  );

  if (response.statusCode == 200) {
    final data = jsonDecode(response.body);
    return data['id']; // сервер должен возвращать client ID
  } else {
    print("Ошибка добавления клиента: ${response.body}");
    return null;
  }
}

Future<bool> makeOrder(int clientId, String carVin, String phone) async {
  final response = await http.post(
    Uri.parse('http://localhost:8081/MakeOrder'),
    headers: {'Content-Type': 'application/json'},
    body: jsonEncode({
      "client_id": clientId,
      "vin": carVin,
      "phone": phone,
    }),
  );

  if (response.statusCode != 200) {
    print("MakeOrder failed: ${response.body}");
  }

  return response.statusCode == 200;
}

Future<List<Order>> ShowBoughtCar() async {
  final response = await http.get(
    Uri.parse('http://localhost:8080/BoughtCatalog'),
    headers: {'Content-Type': 'application/json'},
  );

  if (response.statusCode != 200) {
    throw Exception("Failed to load orders");
  }

  final List<dynamic> jsonData = json.decode(response.body);
  return jsonData.map((e) => Order.fromJson(e)).toList();
}