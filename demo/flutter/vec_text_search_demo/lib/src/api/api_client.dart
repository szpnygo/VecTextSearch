import 'dart:convert';
import 'package:http/http.dart' as http;

class ApiClient {
  static const String _baseUrl = 'http://127.0.0.1:8000';

  static Future<List<Map<String, dynamic>>> searchSimilarTexts(
      String content) async {
    final response = await http.post(
      Uri.parse('$_baseUrl/search-similar-texts'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({'content': content}),
    );

    if (response.statusCode == 200) {
      final List<dynamic> responseData = json.decode(response.body);
      return responseData
          .map((entry) => Map<String, dynamic>.from(entry))
          .toList();
    } else {
      throw Exception('Failed to search for similar texts');
    }
  }
}
