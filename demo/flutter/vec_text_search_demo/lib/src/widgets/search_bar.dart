import 'package:flutter/material.dart';
import 'package:vec_text_search_demo/src/api/api_client.dart';
import 'package:vec_text_search_demo/src/widgets/search_result_list.dart';

class SearchBar extends StatefulWidget {
  final Function(String) onSearch;
  SearchBar({required this.onSearch});
  @override
  _SearchBarState createState() => _SearchBarState();
}

class _SearchBarState extends State<SearchBar> {
  String _searchQuery = '';

  void _updateSearchQuery(String newQuery) {
    setState(() {
      _searchQuery = newQuery;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: const EdgeInsets.all(16.0),
      child: Column(
        children: [
          TextField(
            onChanged: _updateSearchQuery,
            onSubmitted: (_) => widget.onSearch(_searchQuery),
            decoration: InputDecoration(
              labelText: 'Search',
              hintText: 'Enter text to search for similar texts',
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(12.0),
              ),
              focusedBorder: OutlineInputBorder(
                borderSide: BorderSide(color: Theme.of(context).primaryColor),
                borderRadius: BorderRadius.circular(12.0),
              ),
              // ...
            ),
          ),
        ],
      ),
    );
  }
}
