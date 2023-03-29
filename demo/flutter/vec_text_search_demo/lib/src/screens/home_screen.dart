import 'package:flutter/material.dart';
import 'package:vec_text_search_demo/src/api/api_client.dart';
import 'package:vec_text_search_demo/src/widgets/search_bar.dart';
import 'package:vec_text_search_demo/src/widgets/search_result_list.dart';

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  List<Map<String, dynamic>> _searchResults = [];
  bool _isLoading = false;

  Future<void> _search(String query) async {
    setState(() {
      _isLoading = true;
    });

    try {
      final searchResults = await ApiClient.searchSimilarTexts(query);
      setState(() {
        _searchResults = searchResults;
      });
    } catch (e) {
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Error: ${e.toString()}')),
      );
    } finally {
      setState(() {
        _isLoading = false;
      });
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('VecTextSearch Demo'),
      ),
      body: Stack(
        children: [
          Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              SearchBar(onSearch: _search),
              SearchResultList(searchResults: _searchResults),
            ],
          ),
          if (_isLoading)
            Container(
              color: Colors.black.withOpacity(0.5),
              child: Center(
                child: CircularProgressIndicator(),
              ),
            ),
        ],
      ),
    );
  }
}
