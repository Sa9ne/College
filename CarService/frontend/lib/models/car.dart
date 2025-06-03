class Car {
  final String vin;
  final String brand;
  final String name;
  final int power;
  final String color;
  final int price;
  final int mileage;
  final int year;
  final String condition;
  final int owners;
  final bool boughtStatus;
  final String img;

  Car({
    required this.vin,
    required this.brand,
    required this.name,
    required this.power,
    required this.color,
    required this.price,
    required this.mileage,
    required this.year,
    required this.condition,
    required this.owners,
    required this.boughtStatus,
    required this.img,
  });

  factory Car.fromJson(Map<String, dynamic> json) {
    return Car(
      vin: json['vin'],
      brand: json['brand'],
      name: json['name'],
      power: json['power'],
      color: json['color'],
      price: json['price'],
      mileage: json['mileage'],
      year: json['year'],
      condition: json['condition'],
      owners: json['owner'],
      boughtStatus: json['boughtStatus'],
      img: json['img'],
    );
  }
}
