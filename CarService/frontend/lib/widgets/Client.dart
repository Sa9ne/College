import 'package:flutter/material.dart';
import 'package:frontend/services/api_service.dart';
import 'package:intl/intl.dart';

void showClientFormDialog(BuildContext context, String selectedCarId) {
  final _formKey = GlobalKey<FormState>();

  final nameController = TextEditingController();
  final surnameController = TextEditingController();
  final phoneController = TextEditingController();
  DateTime? selectedBirthdate;

  showDialog(
    context: context,
    builder: (BuildContext context) {
      return AlertDialog(
        shape: RoundedRectangleBorder(borderRadius: BorderRadius.circular(12)),
        title: const Text('Введите информацию о себе'),
        content: StatefulBuilder(
          builder: (context, setState) {
            return Form(
              key: _formKey,
              child: SingleChildScrollView(
                child: Column(
                  mainAxisSize: MainAxisSize.min,
                  children: [
                    TextFormField(
                      controller: nameController,
                      decoration: const InputDecoration(labelText: 'Имя'),
                      validator: (value) => value == null || value.isEmpty ? 'Введите имя' : null,
                    ),
                    TextFormField(
                      controller: surnameController,
                      decoration: const InputDecoration(labelText: 'Фамилия'),
                      validator: (value) => value == null || value.isEmpty ? 'Введите фамилию' : null,
                    ),
                    TextFormField(
                      controller: phoneController,
                      decoration: const InputDecoration(labelText: 'Телефон'),
                      validator: (value) => value == null || value.isEmpty ? 'Введите телефон' : null,
                      keyboardType: TextInputType.phone,
                    ),
                    const SizedBox(height: 12),
                    Row(
                      children: [
                        Expanded(
                          child: Text(
                            selectedBirthdate == null
                                ? 'Дата рождения не выбрана'
                                : DateFormat('dd.MM.yyyy').format(selectedBirthdate!),
                          ),
                        ),
                        TextButton(
                          onPressed: () async {
                            final picked = await showDatePicker(
                              context: context,
                              initialDate: DateTime(2000),
                              firstDate: DateTime(1900),
                              lastDate: DateTime.now(),
                            );
                            if (picked != null) {
                              setState(() {
                                selectedBirthdate = picked;
                              });
                            }
                          },
                          child: const Text('Выбрать дату'),
                        )
                      ],
                    )
                  ],
                ),
              ),
            );
          },
        ),
        actions: [
          TextButton(
            onPressed: () => Navigator.of(context).pop(),
            child: const Text('Отмена'),
          ),
          ElevatedButton(
            onPressed: () async {
              if (_formKey.currentState!.validate() && selectedBirthdate != null) {
                final clientData = {
                  "name": nameController.text,
                  "surname": surnameController.text,
                  "birthdate": selectedBirthdate!.toIso8601String() + "Z",
                  "phone": phoneController.text,
                };

                final clientId = await addClient(clientData);

                if (clientId != null) {
                  final success = await makeOrder(clientId, selectedCarId, phoneController.text);
                  if (success) {
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text('Заказ оформлен!')),
                    );
                    Navigator.pop(context);
                  } else {
                    ScaffoldMessenger.of(context).showSnackBar(
                      const SnackBar(content: Text('Ошибка оформления заказа')),
                    );
                  }
                } else {
                  ScaffoldMessenger.of(context).showSnackBar(
                    const SnackBar(content: Text('Ошибка добавления клиента')),
                  );
                }
              } else if (selectedBirthdate == null) {
                ScaffoldMessenger.of(context).showSnackBar(
                  const SnackBar(content: Text('Выберите дату рождения')),
                );
              }
            },
            child: const Text('Сохранить'),
          )
        ],
      );
    },
  );
}
